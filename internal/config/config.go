package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Port            string
	Environment     string
	LogLevelName    string
	ShutdownTimeout time.Duration
}

func Load() (Config, error) {
	cfg := Config{
		Port:            env("APP_PORT", "8080"),
		Environment:     env("APP_ENV", "development"),
		LogLevelName:    strings.ToLower(env("LOG_LEVEL", "info")),
		ShutdownTimeout: 10 * time.Second,
	}
	port, err := strconv.Atoi(cfg.Port)
	if err != nil || port < 1 || port > 65535 {
		return Config{}, fmt.Errorf("APP_PORT must be between 1 and 65535")
	}
	if _, ok := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}[cfg.LogLevelName]; !ok {
		return Config{}, fmt.Errorf("LOG_LEVEL must be debug, info, warn, or error")
	}
	return cfg, nil
}

func (c Config) LogLevel() slog.Level {
	return map[string]slog.Level{"debug": slog.LevelDebug, "warn": slog.LevelWarn, "error": slog.LevelError}[c.LogLevelName]
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && strings.TrimSpace(value) != "" {
		return value
	}
	return fallback
}
