package httpapi

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"sync/atomic"
	"time"
)

type API struct {
	logger   *slog.Logger
	version  string
	requests atomic.Uint64
	errors   atomic.Uint64
	started  time.Time
}

type response struct {
	Data any `json:"data,omitempty"`
	Meta any `json:"meta,omitempty"`
}

func New(logger *slog.Logger, version string) http.Handler {
	a := &API{logger: logger, version: version, started: time.Now().UTC()}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", a.health)
	mux.HandleFunc("GET /ready", a.ready)
	mux.HandleFunc("GET /metrics", a.metrics)
	mux.HandleFunc("GET /api/v1/catalog", a.catalog)
	mux.HandleFunc("POST /api/v1/sessions", a.createSession)
	return a.recover(a.requestID(a.log(a.securityHeaders(mux))))
}

func (a *API) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, response{Data: map[string]string{"status": "healthy", "version": a.version}})
}

func (a *API) ready(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, response{Data: map[string]string{"status": "ready"}})
}

func (a *API) catalog(w http.ResponseWriter, _ *http.Request) {
	items := []map[string]string{
		{"id": "film-001", "title": "Reliability at Dawn", "type": "movie"},
		{"id": "live-001", "title": "Global Live Session", "type": "live"},
	}
	writeJSON(w, http.StatusOK, response{Data: items, Meta: map[string]int{"count": len(items)}})
}

func (a *API) createSession(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ContentID string `json:"content_id"`
	}
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil || input.ContentID == "" {
		a.errors.Add(1)
		writeJSON(w, http.StatusBadRequest, response{Data: map[string]string{"error": "content_id is required"}})
		return
	}
	sessionID := fmt.Sprintf("session-%d", time.Now().UTC().UnixNano())
	w.Header().Set("Location", "/api/v1/sessions/"+sessionID)
	writeJSON(w, http.StatusCreated, response{Data: map[string]string{
		"id": sessionID, "content_id": input.ContentID, "status": "started",
	}})
}

func (a *API) metrics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	fmt.Fprintf(w, "# HELP streaming_api_requests_total Total HTTP requests.\n# TYPE streaming_api_requests_total counter\nstreaming_api_requests_total %d\n", a.requests.Load())
	fmt.Fprintf(w, "# HELP streaming_api_errors_total Total handled API errors.\n# TYPE streaming_api_errors_total counter\nstreaming_api_errors_total %d\n", a.errors.Load())
	fmt.Fprintf(w, "# HELP streaming_api_uptime_seconds Process uptime.\n# TYPE streaming_api_uptime_seconds gauge\nstreaming_api_uptime_seconds %.0f\n", time.Since(a.started).Seconds())
}

func (a *API) requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" || len(id) > 128 {
			id = fmt.Sprintf("req-%d", time.Now().UnixNano())
		}
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r)
	})
}

func (a *API) log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		a.requests.Add(1)
		next.ServeHTTP(w, r)
		a.logger.Info("request completed", "method", r.Method, "path", r.URL.Path, "duration_ms", time.Since(start).Milliseconds())
	})
}

func (a *API) securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		next.ServeHTTP(w, r)
	})
}

func (a *API) recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if value := recover(); value != nil {
				a.errors.Add(1)
				a.logger.Error("panic recovered", "value", value, "stack", string(debug.Stack()))
				writeJSON(w, http.StatusInternalServerError, response{Data: map[string]string{"error": "internal server error"}})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, body response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
