package main

import (
	"github.com/docker/github-actions/internal/command"
	"github.com/docker/github-actions/internal/options"
)

func buildPush(cmd command.Runner) error {
	github, err := options.GetGitHubOptions()
	if err != nil {
		return err
	}

	registry := options.GetRegistry()
	tags := options.GetTags(registry, github)

	build, err := options.GetBuildOptions()
	if err != nil {
		return err
	}
	if err = command.RunBuild(cmd, build, github, tags); err != nil {
		return err
	}

	if options.ShouldPush() {
		login, err := options.GetLoginOptions()
		if err != nil {
			return err
		}
		if login.Username != "" && login.Password != "" {
			if err := command.RunLogin(cmd, login, registry); err != nil {
				return err
			}
		}
	}

	return command.RunPush(cmd, tags)
}
