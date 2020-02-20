TAG ?= latest
STATIC_FLAGS = BUILDKIT_PROGRESS=plain
DOCKER_BUILD = $(STATIC_FLAGS) docker build

all:
	$(DOCKER_BUILD) -t docker/github-actions:$(TAG) .

image:
	$(DOCKER_BUILD) -t docker/github-actions:$(TAG) --build-arg MAKE_TARGET=build .

cli:
	@$(call mkdir,bin)
	$(DOCKER_BUILD) -t github-actions-cli --target=cli --output type=local,dest=./bin/ --build-arg MAKE_TARGET=build .

lint:
	$(DOCKER_BUILD) -t github-actions-lint --build-arg MAKE_TARGET=lint .

test:
	$(DOCKER_BUILD) -t github-actions-test --build-arg MAKE_TARGET=test .
