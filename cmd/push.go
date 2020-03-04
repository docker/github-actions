package main

import (
	"github.com/docker/github-actions/internal/command"
	"github.com/docker/github-actions/internal/options"
)

func push(cmd command.Runner) error {
	github, err := options.GetGitHubOptions()
	if err != nil {
		return err
	}
	tags, err := options.GetTags(options.GetRegistry(), github)
	if err != nil {
		return err
	}
	return command.RunPush(cmd, tags)
}
