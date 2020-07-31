package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/docker/github-actions/internal/options"

	backoff "github.com/cenkalti/backoff/v4"
)

// Runner executes standard commands
type Runner interface {
	// Run executes the given command with arguments
	Run(name string, args ...string) error
}

type execRunner struct{}

// NewRunner returns a new Runner with the stdout and stderr defaulting to os
func NewRunner() Runner {
	return execRunner{}
}

// Run executes the given command with arguments using exec.Command
func (runner execRunner) Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// RunLogin runs a docker login
func RunLogin(cmd Runner, opt options.Login, registry string) error {
	fmt.Println("Logging in to registry", registry)
	args := LoginArgs(opt, registry)
	return cmd.Run("docker", args...)
}

// RunBuild runs a docker build and tags the resulting image
func RunBuild(cmd Runner, opt options.Build, github options.GitHub, tags []string) error {
	fmt.Println("Building image", tags)
	args := BuildArgs(opt, github, tags)
	return cmd.Run("docker", args...)
}

// RunPush runs a docker push for each tag
func RunPush(cmd Runner, tags []string) error {
	fmt.Println("Pushing image", tags)
	for _, tag := range tags {
		args := PushArgs(tag)
		pushBackoff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), options.MaxPushRetries())
		if err := backoff.Retry(func() error {
			return cmd.Run("docker", args...)
		}, pushBackoff); err != nil {
			return err
		}
	}
	return nil
}
