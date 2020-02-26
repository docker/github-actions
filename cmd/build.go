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

	tags := options.GetTags(options.GetServer(), github)

	return command.RunBuild(cmd, o, github, tags)
}
