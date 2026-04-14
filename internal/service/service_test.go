package service

import (
	"context"
	"errors"
	"testing"
	"time"

	ratesv1 "exchange_rate/gen/rates/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type stubClient struct {
	book OrderBook
	err  error
}

func (s stubClient) FetchBook(context.Context) (OrderBook, error) {
	return s.book, s.err
}

type stubRepository struct {
	snapshots []Snapshot
	err       error
}

func (s *stubRepository) SaveRate(_ context.Context, snapshot Snapshot) error {
	if s.err != nil {
		return s.err
	}
	s.snapshots = append(s.snapshots, snapshot)
	return nil
}

func TestGetRatesSuccess(t *testing.T) {
	t.Parallel()

	repo := &stubRepository{}
	svc := New(stubClient{
		book: OrderBook{
			Asks: []string{"95.10", "95.20"},
			Bids: []string{"94.90", "94.80"},
		},
	}, repo, nil, func() time.Time {
		return time.Unix(1700000000, 0).UTC()
	})

	resp, err := svc.GetRates(context.Background(), &ratesv1.GetRatesRequest{
		Method: ratesv1.CalculationMethod_CALCULATION_METHOD_TOP_N,
		N:      1,
	})
	if err != nil {
		t.Fatalf("GetRates() error = %v", err)
	}
	if resp.Ask != "95.10" || resp.Bid != "94.90" {
		t.Fatalf("unexpected response: %+v", resp)
	}
	if got := resp.RetrievedAt.AsTime(); !got.Equal(timestamppb.New(time.Unix(1700000000, 0).UTC()).AsTime()) {
		t.Fatalf("retrieved_at = %s, want %s", got, time.Unix(1700000000, 0).UTC())
	}
	if len(repo.snapshots) != 1 {
		t.Fatalf("saved snapshots = %d, want 1", len(repo.snapshots))
	}
}

func TestGetRatesClientFailure(t *testing.T) {
	t.Parallel()

	repo := &stubRepository{}
	svc := New(stubClient{err: errors.New("upstream down")}, repo, nil, time.Now)

	if _, err := svc.GetRates(context.Background(), &ratesv1.GetRatesRequest{
		Method: ratesv1.CalculationMethod_CALCULATION_METHOD_TOP_N,
		N:      1,
	}); err == nil {
		t.Fatal("GetRates() error = nil, want non-nil")
	}
	if len(repo.snapshots) != 0 {
		t.Fatalf("saved snapshots = %d, want 0", len(repo.snapshots))
	}
}

func TestGetRatesRepositoryFailure(t *testing.T) {
	t.Parallel()

	svc := New(stubClient{
		book: OrderBook{
			Asks: []string{"95.10"},
			Bids: []string{"94.90"},
		},
	}, &stubRepository{err: errors.New("db down")}, nil, time.Now)

	if _, err := svc.GetRates(context.Background(), &ratesv1.GetRatesRequest{
		Method: ratesv1.CalculationMethod_CALCULATION_METHOD_TOP_N,
		N:      1,
	}); err == nil {
		t.Fatal("GetRates() error = nil, want non-nil")
	}
}
