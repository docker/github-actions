ARG GO_VERSION=1.13.7
ARG ALPINE_VERSION=3.11.3



FROM golang:${GO_VERSION} AS builder

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.6

ARG MAKE_TARGET=all
ENV CGO_ENABLED=0
WORKDIR /src

COPY . .

RUN make ${MAKE_TARGET}



FROM alpine:${ALPINE_VERSION}

COPY --from=builder /src/github-actions /github-actions

ENTRYPOINT ["/github-actions"]