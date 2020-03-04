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

	tags, err := options.GetTags(options.GetRegistry(), github)
	if err != nil {
		return err
	}

	return command.RunBuild(cmd, o, github, tags)
}
