APP_BIN := app
GO_IMAGE := golang:1.25
GO_CACHE_DIR := .cache/go-build
GO_MOD_CACHE_DIR := .cache/gomod
GO_PATH_DIR := .cache/gopath
DOCKER_GO := docker run --rm -u $$(id -u):$$(id -g) \
	-e GOCACHE=/src/$(GO_CACHE_DIR) \
	-e GOMODCACHE=/src/$(GO_MOD_CACHE_DIR) \
	-e GOPATH=/src/$(GO_PATH_DIR) \
	-v $(CURDIR):/src -w /src

.PHONY: build test docker-build run lint

build:
	mkdir -p $(GO_CACHE_DIR) $(GO_MOD_CACHE_DIR) $(GO_PATH_DIR)
	$(DOCKER_GO) $(GO_IMAGE) go build -mod=vendor -o $(APP_BIN) ./cmd/app

test:
	mkdir -p $(GO_CACHE_DIR) $(GO_MOD_CACHE_DIR) $(GO_PATH_DIR)
	$(DOCKER_GO) $(GO_IMAGE) go test -mod=vendor ./...

docker-build:
	docker build -t exchange-rate-service .

run:
	./$(APP_BIN)

lint:
	docker run --rm -u $$(id -u):$$(id -g) -v $(CURDIR):/app -w /app golangci/golangci-lint:v2.1.6 golangci-lint run
