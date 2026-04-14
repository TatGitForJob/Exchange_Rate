package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	defaultGRPCPort        = 50051
	defaultDatabaseDSN     = "postgres://postgres:postgres@localhost:5432/exchange_rate?sslmode=disable"
	defaultGrinexTimeout   = 5 * time.Second
	defaultShutdownTimeout = 10 * time.Second
)

// Config stores all runtime configuration for the application.
type Config struct {
	GRPCPort        int
	DatabaseDSN     string
	GrinexTimeout   time.Duration
	ShutdownTimeout time.Duration
	OTLPEndpoint    string
}

// Load resolves configuration from flags and environment variables.
func Load(args []string) (Config, error) {
	fs := flag.NewFlagSet("exchange-rate", flag.ContinueOnError)

	cfg := Config{}
	fs.IntVar(&cfg.GRPCPort, "grpc-port", envInt("EXCHANGE_GRPC_PORT", defaultGRPCPort), "gRPC listen port")
	fs.StringVar(&cfg.DatabaseDSN, "database-dsn", envString("EXCHANGE_DATABASE_DSN", defaultDatabaseDSN), "PostgreSQL DSN")
	fs.DurationVar(&cfg.GrinexTimeout, "grinex-timeout", envDuration("EXCHANGE_GRINEX_TIMEOUT", defaultGrinexTimeout), "Grinex request timeout")
	fs.DurationVar(&cfg.ShutdownTimeout, "shutdown-timeout", envDuration("EXCHANGE_SHUTDOWN_TIMEOUT", defaultShutdownTimeout), "Graceful shutdown timeout")
	fs.StringVar(&cfg.OTLPEndpoint, "otlp-endpoint", envString("EXCHANGE_OTLP_ENDPOINT", ""), "OTLP gRPC endpoint")

	if err := fs.Parse(args[1:]); err != nil {
		return Config{}, err
	}

	if cfg.GRPCPort <= 0 || cfg.GRPCPort > 65535 {
		return Config{}, fmt.Errorf("grpc port must be between 1 and 65535")
	}
	if cfg.DatabaseDSN == "" {
		return Config{}, fmt.Errorf("database dsn must not be empty")
	}
	return cfg, nil
}

func envString(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func envDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return duration
}

func envInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
