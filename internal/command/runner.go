package command

import (
	"os"
	"os/exec"
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
