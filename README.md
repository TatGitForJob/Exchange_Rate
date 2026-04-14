# Exchange Rate Service

gRPC service that fetches the USDT market depth from CEX(Grinex), calculates `ask` and `bid` prices, stores every successful snapshot in PostgreSQL, exposes gRPC health checks, and supports graceful shutdown.

## Requirements

- Docker and Docker Compose

## Configuration

Flags override environment variables, environment variables override defaults.

| Flag | Environment variable | Default |
| --- | --- | --- |
| `-grpc-port` | `EXCHANGE_GRPC_PORT` | `50051` |
| `-database-dsn` | `EXCHANGE_DATABASE_DSN` | `postgres://postgres:postgres@localhost:5432/exchange_rate?sslmode=disable` |
| `-grinex-timeout` | `EXCHANGE_GRINEX_TIMEOUT` | `5s` |
| `-shutdown-timeout` | `EXCHANGE_SHUTDOWN_TIMEOUT` | `10s` |
| `-otlp-endpoint` | `EXCHANGE_OTLP_ENDPOINT` | empty |

## Build and Run

Build the binary:

```bash
make build
```

Run unit tests:

```bash
make test
```

Run the application in Docker after PostgreSQL is ready:

```bash
docker-compose run --rm app ./app
```

## gRPC API

Service definition is stored in `api/proto/rates/v1/rates.proto`.

`GetRates` accepts:
- `method = CALCULATION_METHOD_TOP_N` with `n >= 1`
- `method = CALCULATION_METHOD_AVG_N_M` with `1 <= n <= m`

Positions are 1-based and the average range is inclusive.

Example `grpcurl` requests:

```bash
grpcurl -plaintext -d '{}' localhost:50051 grpc.health.v1.Health/Check
```
```bash
grpcurl -plaintext -d '{"method":"CALCULATION_METHOD_TOP_N","n":1}' localhost:50051 rates.v1.RatesService/GetRates
```

The service also registers the standard `grpc.health.v1.Health` service.

## Migrations

Migrations are embedded into the binary from the `migrations/` directory and are applied automatically on startup before the gRPC server starts accepting traffic.

## Observability

OpenTelemetry tracing is enabled for inbound gRPC requests, outbound HTTP requests to Grinex, and PostgreSQL operations. `EXCHANGE_OTLP_ENDPOINT` accepts either `host:port` or a full collector URL such as `http://otel-collector:4317`.

## Lint

```bash
make lint
```
