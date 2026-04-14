FROM golang:1.25 AS builder

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
COPY vendor ./vendor

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o /out/app ./cmd/app

FROM scratch

WORKDIR /app

COPY --from=builder /out/app /app/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 50051

ENTRYPOINT ["/app/app"]
