default: build lint test-unit

all: default test-e2e

build:
	@$(call mkdir,bin)
	go build -o bin/github-actions ./cmd

lint:
	golangci-lint run --config golangci.yml ./...

test: test-unit test-e2e

test-unit:
	go test $(go list ./... | grep -v /e2e)

test-e2e: build
	go test ./e2e/...
