package options

import "github.com/caarlos0/env/v6"

// GitHub contains the parsed common github actions environment variables
// See https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables
type GitHub struct {
	RunInActions bool   `env:"GITHUB_ACTIONS"`
	Workflow     string `env:"GITHUB_WORKFLOW"`
	RunID        string `env:"GITHUB_RUN_ID"`
	RunNumber    string `env:"GITHUB_RUN_NUMBER"`
	Action       string `env:"GITHUB_ACTION"`
	Actor        string `env:"GITHUB_ACTOR"`
	Repository   string `env:"GITHUB_REPOSITORY"`
	EventName    string `env:"GITHUB_EVENT_NAME"`
	Sha          string `env:"GITHUB_SHA"`
	Ref          string `env:"GITHUB_REF"`
}

// GetGitHubOptions gets the common github actions environment variables
func GetGitHubOptions() (GitHub, error) {
	var github GitHub
	err := env.Parse(&github)
	return github, err
}
