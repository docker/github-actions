package main

import (
	"fmt"
	"os"

	"github.com/docker/github-actions/internal/command"
	"github.com/docker/github-actions/internal/options"
	commandLine "github.com/urfave/cli/v2"
)

func main() {
	_, err := options.GetGitHubOptions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := command.NewRunner()

	app := &commandLine.App{
		Name:  "docker github actions",
		Usage: "Used in GitHub Actions to run Docker workflows",
		Commands: []*commandLine.Command{
			{
				Name:        "login",
				Description: "Logs into a docker server",
				Action: func(c *commandLine.Context) error {
					return login(cmd)
				},
			},
			{
				Name:        "build",
				Description: "Builds a docker image",
				Action: func(c *commandLine.Context) error {
					return build(cmd)
				},
			},
			{
				Name:        "push",
				Description: "Pushes a docker image",
				Action: func(c *commandLine.Context) error {
					return push(cmd)
				},
			},
		},
	}

	if err = app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
