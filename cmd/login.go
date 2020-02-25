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

	return runLogin(cmd, o, options.GetServer())
}

func runLogin(cmd command.Runner, opt options.Login, server string) error {
	args := command.LoginArgs(opt, server)
	return cmd.Run("docker", args...)
}
