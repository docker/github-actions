package options

import (
	"os"
	"strings"

	"github.com/caarlos0/env/v6"
)

// GitReferenceType is the type of the git ref
type GitReferenceType int

// Get reference types
const (
	GitRefUnknown GitReferenceType = iota
	GitRefHead
	GitRefPullRequest
	GitRefTag
)

// GitReference contains the type and name of the git ref
type GitReference struct {
	Type GitReferenceType
	Name string
}

func parseGitRef(ref string) GitReference {
	if split := strings.SplitN(ref, "/", 3); len(split) == 3 {
		name := split[2]
		switch split[1] {
		case "heads":
			return GitReference{GitRefHead, name}
		case "pulls":
			return GitReference{GitRefPullRequest, name}
		case "tags":
			return GitReference{GitRefTag, name}
		}
	}
	return GitReference{GitRefUnknown, ""}
}

// GitHub contains the parsed common github actions environment variables
// See https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables
type GitHub struct {
	RunInActions bool   `env:"GITHUB_ACTIONS"`
	Workflow     string `env:"GITHUB_WORKFLOW"`
	Action       string `env:"GITHUB_ACTION"`
	Actor        string `env:"GITHUB_ACTOR"`
	Repository   string `env:"GITHUB_REPOSITORY"`
	EventName    string `env:"GITHUB_EVENT_NAME"`
	Sha          string `env:"GITHUB_SHA"`
	Reference    GitReference
}

// GetGitHubOptions gets the common github actions environment variables
func GetGitHubOptions() (GitHub, error) {
	var github GitHub
	err := env.Parse(&github)
	if err != nil {
		return github, err
	}
	github.Reference = parseGitRef(os.Getenv("GITHUB_REF"))
	return github, nil
}
