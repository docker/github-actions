TAG ?= latest
STATIC_FLAGS = BUILDKIT_PROGRESS=plain
DOCKER_BUILD = $(STATIC_FLAGS) docker build

ROOT_DIR = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

all:
	$(DOCKER_BUILD) -t docker/github-actions:$(TAG) .

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
	docker run --rm --network="host" -e "E2E_HOST_PATH=$(ROOT_DIR)/e2e" -v /var/run/docker.sock:/var/run/docker.sock github-actions-test-e2e make test-e2e