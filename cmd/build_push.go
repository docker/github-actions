package main

import (
	"fmt"

	"github.com/docker/github-actions/internal/command"
	"github.com/docker/github-actions/internal/options"
)

func buildPush(cmd command.Runner) error {
	registry := options.GetRegistry()
	login, err := options.GetLoginOptions()
	if err != nil {
		return err
	}
	if login.Username != "" && login.Password != "" {
		if err := command.RunLogin(cmd, login, registry); err != nil {
			return err
		}
	}

	github, err := options.GetGitHubOptions()
	if err != nil {
		return err
	}

	tags, err := options.GetTags(registry, github)
	if err != nil {
		return err
	}

	build, err := options.GetBuildOptions()
	if err != nil {
		return err
	}
	if err = command.RunBuild(cmd, build, github, tags); err != nil {
		return err
	}

	if shouldPush, err := options.ShouldPush(); err != nil {
		return err
	} else if shouldPush {
		return command.RunPush(cmd, tags)
	}

	fmt.Println("Skipping push")
	return nil
}
