package command

import (
	"fmt"

	"github.com/docker/github-actions/internal/options"
)

// LoginArgs converts login options into the cli arguments used to call `docker login`
func LoginArgs(o options.Login) []string {
	args := []string{"login", "--username", o.Username, "--password", o.Password}
	if o.Server != "" {
		args = append(args, o.Server)
	}
	return args
}

// BuildArgs converts build options into the cli arguments used to call `docker build`
func BuildArgs(o options.Build, github options.GitHub) []string {
	args := []string{"build"}

	for _, tag := range options.GetTags(o, github) {
		t := fmt.Sprintf("%s:%s", o.Repository, tag)
		if o.Server != "" {
			t = fmt.Sprintf("%s/%s", o.Server, t)
		}
		args = append(args, "-t", t)
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
