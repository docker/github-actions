# github-actions
The core code base for Docker's GitHub Actions (https://github.com/features/actions). This code is used to build the docker/github-actions image that provides the functionality used by the published Docker GitHub Actions


## Building github-actions
The code is written in Go v1.13 with `go mod`. It can be built locally using the `Makefile` or in docker using the `docker.Makefile`.

`make -f docker.Makefile` will build the code, check the linting using golangci-lint, run the go tests, and build the image with a tag of docker/github-actions:latest

`make -f docker.Makefile TAG=foo` will build the code, check the linting using golangci-lint, run the go tests, and build the image with a tag of docker/github-actions:foo

`make -f docker.Makefile image` will build the github-actions image without a tag and without running test or lint checking

`make -f docker.Makefile cli` will build the cli and copy it to `./bin/github-actions`

`make -f docker.Makefile test` will run the go tests
