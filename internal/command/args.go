package command

import "github.com/docker/github-actions/internal/options"

func LoginArgs(o options.Login) []string {
	args := []string{"login", "--username", o.Username, "--password", o.Password}
	if o.Server != "" {
		args = append(args, o.Server)
	}
	return args
}
