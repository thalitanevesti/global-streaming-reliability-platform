# Global Streaming Reliability Platform

A production-oriented portfolio project for designing, securing, observing, and operating a globally distributed streaming control plane.

## Milestone 1

- Go 1.27 API with graceful shutdown and structured logs
- Liveness, readiness, and Prometheus-compatible metrics
- Versioned catalog and streaming-session endpoints
- Tests with race detection and coverage
- Minimal non-root container with a read-only filesystem
- Resource limits, dropped Linux capabilities, and no-new-privileges
- CI checks for tests, vet, formatting, and compilation

## Run locally

```bash
docker compose up --build -d
curl http://localhost:18080/health
curl http://localhost:18080/api/v1/catalog
curl -X POST http://localhost:18080/api/v1/sessions \
  -H 'Content-Type: application/json' \
  -d '{"content_id":"film-001"}'
```

Stop the environment:

```bash
docker compose down --remove-orphans
```

## Roadmap

1. PostgreSQL, Redis, migrations, and integration tests
2. OpenTelemetry, Prometheus, Grafana, and initial SLOs
3. Kubernetes manifests and local `kind` environment
4. Secure CI/CD with software supply-chain controls
5. Modular AWS infrastructure with Terraform
6. Load tests, failure injection, runbooks, and postmortems

No AWS resources are provisioned by this repository unless an operator explicitly runs an approved apply workflow.

## Documentation

- [Architecture](docs/architecture.md)
- [ADR 0001: Use Go](docs/adr/0001-use-go-for-control-plane-api.md)

## Author

Thalita S. Neves — [GitHub](https://github.com/thalitanevesti) · [LinkedIn](https://linkedin.com/in/thalitanevesti)

