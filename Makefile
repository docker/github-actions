all: build lint test

build:
	go build -o github-actions ./cmd

lint:
	golangci-lint run --config golangci.yml ./...

test:
	go test ./...
