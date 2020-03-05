package main

import (
	"fmt"
	"os"

	"github.com/docker/github-actions/internal/command"
	"github.com/spf13/cobra"
)

func main() {
	runner := command.NewRunner()

	rootCmd := &cobra.Command{
		Use:   "github-actions",
		Short: "Used in GitHub Actions to run Docker workflows",
	}
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "login",
			Short: "Logs into a docker registry",
			RunE: func(cmd *cobra.Command, args []string) error {
				return login(runner)
			},
		},
		&cobra.Command{
			Use:   "build",
			Short: "Builds a docker image",
			RunE: func(cmd *cobra.Command, args []string) error {
				return build(runner)
			},
		},
		&cobra.Command{
			Use:   "push",
			Short: "Pushes a docker image",
			RunE: func(cmd *cobra.Command, args []string) error {
				return push(runner)
			},
		},
		&cobra.Command{
			Use:   "build-push",
			Short: "Builds and pushes a docker image to a registry, logging in if necessary",
			RunE: func(cmd *cobra.Command, args []string) error {
				return buildPush(runner)
			},
		},
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
