package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	ratesv1 "exchange_rate/gen/rates/v1"
	"exchange_rate/internal/calculator"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// OrderBook contains exchange prices split by side.
type OrderBook struct {
	Asks []string
	Bids []string
}

// Snapshot contains data persisted for each successful request.
type Snapshot struct {
	Method      string
	N           uint32
	M           *uint32
	Ask         string
	Bid         string
	RetrievedAt time.Time
}

// Repository persists calculated rate snapshots.
type Repository interface {
	SaveRate(ctx context.Context, snapshot Snapshot) error
}

type marketClient interface {
	FetchBook(ctx context.Context) (OrderBook, error)
}

// Service implements the rates gRPC service.
type Service struct {
	ratesv1.UnimplementedRatesServiceServer

	client marketClient
	repo   Repository
	tracer trace.Tracer
	now    func() time.Time
}

// New constructs a new service instance.
func New(client marketClient, repo Repository, tracerProvider trace.TracerProvider, now func() time.Time) *Service {
	if tracerProvider == nil {
		tracerProvider = otel.GetTracerProvider()
	}
	if now == nil {
		now = time.Now
	}

	return &Service{
		client: client,
		repo:   repo,
		tracer: tracerProvider.Tracer("exchange_rate/internal/service"),
		now:    now,
	}
}

// GetRates fetches, calculates, stores, and returns the current rate snapshot.
func (s *Service) GetRates(ctx context.Context, req *ratesv1.GetRatesRequest) (*ratesv1.GetRatesResponse, error) {
	ctx, span := s.tracer.Start(ctx, "Service.GetRates")
	defer span.End()

	method, snapshotMethod, err := mapMethod(req.GetMethod())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := validateRequest(req, method); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	book, err := s.client.FetchBook(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "fetch order book: %v", err)
	}

	ask, err := calculator.Calculate(book.Asks, method, req.GetN(), req.GetM())
	if err != nil {
		return nil, translateCalculationError(err)
	}
	bid, err := calculator.Calculate(book.Bids, method, req.GetN(), req.GetM())
	if err != nil {
		return nil, translateCalculationError(err)
	}

	retrievedAt := s.now().UTC()
	snapshot := Snapshot{
		Method:      snapshotMethod,
		N:           req.GetN(),
		Ask:         ask,
		Bid:         bid,
		RetrievedAt: retrievedAt,
	}
	if method == calculator.MethodAvgNM {
		m := req.GetM()
		snapshot.M = &m
	}

	if err := s.repo.SaveRate(ctx, snapshot); err != nil {
		return nil, status.Errorf(codes.Internal, "save snapshot: %v", err)
	}

	return &ratesv1.GetRatesResponse{
		Ask:         ask,
		Bid:         bid,
		RetrievedAt: timestamppb.New(retrievedAt),
	}, nil
}

func mapMethod(method ratesv1.CalculationMethod) (calculator.Method, string, error) {
	switch method {
	case ratesv1.CalculationMethod_CALCULATION_METHOD_TOP_N:
		return calculator.MethodTopN, string(calculator.MethodTopN), nil
	case ratesv1.CalculationMethod_CALCULATION_METHOD_AVG_N_M:
		return calculator.MethodAvgNM, string(calculator.MethodAvgNM), nil
	default:
		return "", "", fmt.Errorf("unsupported calculation method %s", method.String())
	}
}

func validateRequest(req *ratesv1.GetRatesRequest, method calculator.Method) error {
	switch method {
	case calculator.MethodTopN:
		if req.GetN() == 0 {
			return errors.New("n must be greater than zero")
		}
	case calculator.MethodAvgNM:
		if req.GetN() == 0 || req.GetM() == 0 || req.GetN() > req.GetM() {
			return errors.New("avgNM requires 1 <= n <= m")
		}
	}
	return nil
}

func translateCalculationError(err error) error {
	if errors.Is(err, context.Canceled) {
		return status.Error(codes.Canceled, err.Error())
	}
	return status.Errorf(codes.InvalidArgument, "calculate rate: %v", err)
}
