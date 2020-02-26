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
	tags := options.GetTags(options.GetServer(), github)
	return command.RunPush(cmd, tags)
}
