package main

import (
	"github.com/docker/github-actions/internal/command"
	"github.com/docker/github-actions/internal/options"
)

func login(cmd command.Runner) error {
	o, err := options.GetLoginOptions()
	if err != nil {
		return err
	}

	args := command.LoginArgs(o)
	return cmd.Run("docker", args...)
}
