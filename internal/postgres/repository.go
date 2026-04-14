package postgres

import (
	"context"
	"fmt"

	"exchange_rate/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Repository stores snapshots in PostgreSQL.
type Repository struct {
	pool   *pgxpool.Pool
	tracer trace.Tracer
}

// NewRepository creates a PostgreSQL-backed repository.
func NewRepository(pool *pgxpool.Pool, tracerProvider trace.TracerProvider) *Repository {
	if tracerProvider == nil {
		tracerProvider = otel.GetTracerProvider()
	}
	return &Repository{
		pool:   pool,
		tracer: tracerProvider.Tracer("exchange_rate/internal/postgres"),
	}
}

// SaveRate inserts a newly calculated snapshot.
func (r *Repository) SaveRate(ctx context.Context, snapshot service.Snapshot) error {
	ctx, span := r.tracer.Start(ctx, "Repository.SaveRate")
	defer span.End()

	_, err := r.pool.Exec(ctx, `
		INSERT INTO rate_snapshots (method, n, m, ask, bid, retrieved_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, snapshot.Method, snapshot.N, snapshot.M, snapshot.Ask, snapshot.Bid, snapshot.RetrievedAt)
	if err != nil {
		return fmt.Errorf("insert rate snapshot: %w", err)
	}
	return nil
}
