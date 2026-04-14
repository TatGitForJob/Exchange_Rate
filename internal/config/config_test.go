package config

import (
	"testing"
	"time"
)

func TestLoadFromEnv(t *testing.T) {
	t.Setenv("EXCHANGE_GRPC_PORT", "9000")
	t.Setenv("EXCHANGE_DATABASE_DSN", "postgres://test")
	t.Setenv("EXCHANGE_GRINEX_TIMEOUT", "3s")

	cfg, err := Load([]string{"app"})
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.GRPCPort != 9000 {
		t.Fatalf("GRPCPort = %d, want %d", cfg.GRPCPort, 9000)
	}
	if cfg.DatabaseDSN != "postgres://test" {
		t.Fatalf("DatabaseDSN = %q, want %q", cfg.DatabaseDSN, "postgres://test")
	}
	if cfg.GrinexTimeout != 3*time.Second {
		t.Fatalf("GrinexTimeout = %s, want %s", cfg.GrinexTimeout, 3*time.Second)
	}
}

func TestFlagsOverrideEnv(t *testing.T) {
	t.Setenv("EXCHANGE_GRPC_PORT", "9000")

	cfg, err := Load([]string{"app", "-grpc-port", "9100"})
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.GRPCPort != 9100 {
		t.Fatalf("GRPCPort = %d, want %d", cfg.GRPCPort, 9100)
	}
}
