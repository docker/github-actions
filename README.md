# github-actions
The core code base for Docker's GitHub Actions (https://github.com/features/actions). This code is used to build the docker/github-actions image that provides the functionality used by the published Docker GitHub Actions.

`github-actions` runs a command line tool that shells out to docker to perform the various functions. Parameters are supplied to `github-actions` using environment variables in the form described by the GitHub Actions documentation. `github-actions` uses some of the default GitHub Actions environment variables as described in the individual commands section.

## Commands

Commands can be called using `docker run docker/github-actions {command}`

### login

Does a `docker login` using the supplied username and password. Will default to Docker Hub but can be supplied a server address to login to a third-party registry as required.

#### inputs

|Environment Variable|Required|Description|
|---|---|---|
|INPUT_USERNAME|yes|Username to login with|
|INPUT_PASSWORD|yes|Password to login with|
|INPUT_REGISTRY|no|Registry server to login to. Defaults to Docker Hub|

### build

Builds and tags a docker image.

#### inputs

|Environment Variable|Required|Description|
|---|---|---|
|INPUT_PATH|yes|Path to build from|
|INPUT_DOCKERFILE|no|Path to Dockerfile|
|INPUT_ADD_GIT_LABELS|no|Adds git labels (see below)|
|INPUT_TARGET|no|Target build stage to build|
|INPUT_BUILD_ARGS|no|Comma-delimited list of build-args|
|INPUT_LABELS|no|Comma-delimited list of labels|

See the tagging section for information on tag inputs

##### Git labels

When `INPUT_ADD_GIT_LABELS` is `true` labels are automatically added to the image that contain data about the current state of the git repo:

|Label|Description|
|---|---|
|com.docker.github-actions-actor|The username of the user that kicked off this run of the actions (e.g. the user that did the git push)|
|com.docker.github-actions-sha|The full git sha of this commit|

### push

Pushes a docker image.

#### inputs

See the tagging section for information on tag inputs


### build-push

Builds, logs in, and pushes a docker image.

#### inputs

Same as the login and build commands with the addition of

|Environment Variable|Required|Description|
|---|---|---|
|INPUT_PUSH|no|Will push the image if true|


## Tagging

Tagging of images can be set manually, left to `github-actions` to automate, or a combination of the both.

There are 4 input variables used for tagging

|Environment Variable|Required|Description|
|---|---|---|
|INPUT_REGISTRY|no|Registry server to tag with|
|INPUT_REPOSITORY|yes|Repository to tag with|
|INPUT_TAGS|no|Hard coded comma-delimited list of tags|
|INPUT_TAG_WITH_REF|no|If true then `github-actions` will add tags depending on the git ref automatically as described below|
|INPUT_TAG_WITH_SHA|no|If true then `github-actions` will add a tag in the form `sha-{git-short-sha}`|

If `INPUT_REGISTRY` is set then all tags are prefixed with `{INPUT_REGISTRY}/{INPUT_REPOSITORY}:`.
If not then all tags are prefixed with `{INPUT_REPOSITORY}:`

Auto tags depend on the git reference that the run is associated with. The reference is passed to `github-actions` using the GitHub actions `GITHUB_REF` enviroment variable.

If the reference is `refs/heads/{branch-name}` then the tag `{branch-name}` is added. For the master branch the `{branch-name}` is replaced with `latest`.

If the reference is `refs/pull/{pr}` then the tag `pr-{pr}` is added.

If the reference is `refs/tags/{tag-name}` then the tag `{tag-name}` is added.

Any `/` in the auto tags are replaced with `-`.

For example if the environment variables are as follows:

|Variable|Value|
|---|---|
|INPUT_REGISTRY||
|INPUT_REPOSITORY|myorg/myimage|
|INPUT_TAGS|foo,bar|
|INPUT_TAG_WITH_REF|true|
|GITHUB_REF|refs/tags/v0.1|

Then the image will be tagged with:
```
myorg/myimage:foo
myorg/myimage:bar
myorg/myimage:v0.1
```

If the variables are as follows:

|Variable|Value|
|---|---|
|INPUT_REGISTRY|myregistry|
|INPUT_REPOSITORY|myorg/myimage|
|INPUT_TAGS|foo,bar|
|INPUT_TAG_WITH_REF|true|
|INPUT_TAG_WITH_SHA|true|
|GITHUB_REF|refs/heads/master|
|GITHUB_SHA|c6df8c68eb71799f9c9ab4a4a4650d6aabd7e415|

Then the image will be tagged with:
```
myregistry/myorg/myimage:foo
myregistry/myorg/myimage:bar
myregistry/myorg/myimage:lastest
myregistry/myorg/myimage:c6df8c6
```

## Building github-actions
The code is written in Go v1.13 with `go mod`. It can be built locally using the `Makefile` or in docker using the `docker.Makefile`.

`make -f docker.Makefile` will build the code, check the linting using golangci-lint, run the go tests, and build the image with a tag of docker/github-actions:latest

`make -f docker.Makefile image` will build the github-actions image without a tag and without running test or lint checking

`make -f docker.Makefile cli` will build the cli and copy it to `./bin/github-actions`

`make -f docker.Makefile test` will run the unit and e2e tests
