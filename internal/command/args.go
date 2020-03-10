package command

import (
	"github.com/docker/github-actions/internal/options"
)

// LoginArgs converts login options into the cli arguments used to call `docker login`
func LoginArgs(o options.Login, registry string) []string {
	args := []string{"login", "--username", o.Username, "--password", o.Password}
	if registry != "" {
		args = append(args, registry)
	}
	return args
}

// BuildArgs converts build options into the cli arguments used to call `docker build`
func BuildArgs(o options.Build, github options.GitHub, tags []string) []string {
	args := commonBuildArgs(o, github, tags)
	return append(args, o.Path)
}

// PushArgs converts tags into the cli arguments used to call `docker push`
func PushArgs(tag string) []string {
	return []string{"push", tag}
}

// BuildxCreateArgs returns args to create and use a buildx builder
func BuildxCreateArgs() []string {
	return []string{"buildx", "create", "--name", "github-actions", "--use"}
}

// BuildxBuildArgs converts build options into the cli arguments used to call `docker buildx build`
func BuildxBuildArgs(o options.Build, github options.GitHub, tags []string, push bool) []string {
	args := []string{"buildx"}
	args = append(args, commonBuildArgs(o, github, tags)...)
	args = append(args, "--platform", o.Platforms)

	if push {
		args = append(args, "--push")
	} else {
		args = append(args, "--load")
	}

	return append(args, o.Path)
}

func commonBuildArgs(o options.Build, github options.GitHub, tags []string) []string {
	args := []string{"build", "--progress", "plain"}

	for _, tag := range tags {
		args = append(args, "-t", tag)
	}

	for _, label := range options.GetLabels(o, github) {
		args = append(args, "--label", label)
	}

	if o.Dockerfile != "" {
		args = append(args, "--file", o.Dockerfile)
	}

	if o.Target != "" {
		args = append(args, "--target", o.Target)
	}

	if o.AlwaysPull {
		args = append(args, "--pull")
	}

	for _, buildArg := range o.BuildArgs {
		args = append(args, "--build-arg", buildArg)
	}

	return args
}
