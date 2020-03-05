TAG ?= latest
DOCKER_BUILD = docker build --progress=plain

ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
ROOT_DIR := $(subst .\,,$(ROOT_DIR))

all: build test-e2e

build:
	$(DOCKER_BUILD) -t docker/github-actions-default --build-arg MAKE_TARGET=default .

image:
	$(DOCKER_BUILD) -t docker/github-actions:$(TAG) --build-arg MAKE_TARGET=build .

cli:
	@$(call mkdir,bin)
	$(DOCKER_BUILD) -t github-actions-cli --target=cli --output type=local,dest=./bin/ --build-arg MAKE_TARGET=build .

lint:
	$(DOCKER_BUILD) -t github-actions-lint --target=builder --build-arg MAKE_TARGET=lint .

test: test-unit test-e2e

test-unit:
	$(DOCKER_BUILD) -t github-actions-test-unit --target=builder --build-arg MAKE_TARGET=test-unit .

test-e2e:
	$(DOCKER_BUILD) -t github-actions-test-e2e --target e2e --build-arg MAKE_TARGET=build .
	docker run --rm --network="host" -v /var/run/docker.sock:/var/run/docker.sock github-actions-test-e2e make test-e2e
