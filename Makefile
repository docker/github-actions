all: build lint test

build:
	@$(call mkdir,bin)
	go build -o bin/github-actions ./cmd

lint:
	golangci-lint run --config golangci.yml ./...

test:
	go test ./...
