package command

import (
	"github.com/docker/github-actions/internal/options"
)

// LoginArgs converts login options into the cli arguments used to call `docker login`
func LoginArgs(o options.Login, server string) []string {
	args := []string{"login", "--username", o.Username, "--password", o.Password}
	if server != "" {
		args = append(args, server)
	}
	return args
}

// BuildArgs converts build options into the cli arguments used to call `docker build`
func BuildArgs(o options.Build, github options.GitHub, tags []string) []string {
	args := []string{"build"}

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

	if o.Path == "" {
		args = append(args, ".")
	} else {
		args = append(args, o.Path)
	}
	return args
}

// PushArgs converts tags into the cli arguments used to call `docker push`
func PushArgs(tags []string) []string {
	return append([]string{"push"}, tags...)
}
