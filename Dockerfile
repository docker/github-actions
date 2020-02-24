ARG GO_VERSION=1.13.7
ARG GOLANGCI_LINT_VERSION=v1.23.6
ARG DND_VERSION=19.03


# Builds the github-actions binary, checks linting, and runs unit level tests
FROM golang:${GO_VERSION} AS builder

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_LINT_VERSION}

ARG MAKE_TARGET=default
ENV CGO_ENABLED=0
WORKDIR /src

COPY . .

RUN make ${MAKE_TARGET}


# Used to run e2e tests for github-actions
# This image must be run as a container to run the tests
FROM golang:${GO_VERSION} AS e2e
ARG CLI_CHANNEL=stable
ARG CLI_VERSION=19.03.5

RUN apt-get install -y -q --no-install-recommends coreutils util-linux

ENV CGO_ENABLED=0
ENV GITHUB_ACTIONS_BINARY=/github-actions
WORKDIR /tests

RUN curl -fL https://download.docker.com/linux/static/${CLI_CHANNEL}/x86_64/docker-${CLI_VERSION}.tgz | tar xzO docker/docker > /usr/bin/docker && chmod +x /usr/bin/docker

COPY . .
COPY --from=builder /src/bin/github-actions /github-actions


# Used to extract the github-actions binary
FROM scratch AS cli
COPY --from=builder /src/bin/github-actions github-actions


# The github-actions image that is used by published docker github actions
FROM docker:${DND_VERSION}

COPY --from=builder /src/bin/github-actions /github-actions

ENTRYPOINT ["/github-actions"]
