default: build lint test-unit

all: default test-e2e

build:
	@$(call mkdir,bin)
	go build -o bin/github-actions ./cmd

lint:
	golangci-lint run --config golangci.yml ./...

test: test-unit test-e2e

test-unit:
	go test ./cmd/... ./internal/...

test-e2e: build
	docker build --file ./e2e/Dockerfile.registry -t github-actions-registry ./e2e
	go test ./e2e/...
