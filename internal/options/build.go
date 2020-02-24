package options

import (
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v6"
)

const githubLabelPrefix = "com.docker.github-actions"

// Build contains the parsed build action environment variables
type Build struct {
	Path             string `env:"INPUT_PATH"`
	Dockerfile       string `env:"INPUT_DOCKERFILE"`
	SetDefaultLabels bool   `env:"INPUT_SET_DEFAULT_LABELS"`
	Target           string `env:"INPUT_TARGET"`
	AlwaysPull       bool   `env:"INPUT_ALWAYS_PULL"`
	BuildArgs        []string
	Labels           []string
}

// GetBuildOptions gets the login action environment variables
func GetBuildOptions() (Build, error) {
	var build Build
	if err := env.Parse(&build); err != nil {
		return build, err
	}

	if buildArgs := os.Getenv("INPUT_BUILD_ARGS"); buildArgs != "" {
		build.BuildArgs = strings.Split(buildArgs, ",")
	}

	if labels := os.Getenv("INPUT_LABELS"); labels != "" {
		build.Labels = strings.Split(labels, ",")
	}

	return build, nil
}

// GetLabels gets a list of all labels to build the image with including automatic labels created from github vars when SetDefaultLabels is true
func GetLabels(build Build, github GitHub) []string {
	labels := []string{}
	if build.Labels != nil {
		labels = build.Labels
	}

	if !build.SetDefaultLabels {
		return labels
	}

	if github.Actor != "" {
		labels = append(labels, fmt.Sprintf("%s-actor=%s", githubLabelPrefix, github.Actor))
	}

	if github.Sha != "" {
		labels = append(labels, fmt.Sprintf("%s-sha=%s", githubLabelPrefix, github.Sha))
	}

	return labels
}
