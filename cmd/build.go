package main

import (
	"github.com/docker/github-actions/internal/command"
	"github.com/docker/github-actions/internal/options"
)

func build(cmd command.Runner) error {
	o, err := options.GetBuildOptions()
	if err != nil {
		return err
	}

	github, err := options.GetGitHubOptions()
	if err != nil {
		return err
	}

	args := command.BuildArgs(o, github)
	return cmd.Run("docker", args...)
}
