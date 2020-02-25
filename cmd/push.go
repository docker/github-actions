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
	return runPush(cmd, tags)
}

func runPush(cmd command.Runner, tags []string) error {
	for _, tag := range tags {
		args := command.PushArgs(tag)
		if err := cmd.Run("docker", args...); err != nil {
			return err
		}
	}
	return nil
}
