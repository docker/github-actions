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
	Dockerfile       string `env:"INPUT_DOCKERFILE"`
	Server           string `env:"INPUT_SERVER"`
	Repository       string `env:"INPUT_REPOSITORY"`
	SetDefaultTags   bool   `env:"INPUT_SET_DEFAULT_TAGS"`
	SetDefaultLabels bool   `env:"INPUT_SET_DEFAULT_LABELS"`
	Target           string `env:"INPUT_TARGET"`
	AlwaysPull       bool   `env:"INPUT_ALWAYS_PULL"`
	BuildArgs        []string
	Labels           []string
	Tags             []string
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

	if tags := os.Getenv("INPUT_TAGS"); tags != "" {
		build.Tags = strings.Split(tags, ",")
	}

	return build, nil
}

// GetTags gets a list of all tags to build the image with including automatic tags created from github vars when SetDefaultTags is true
func GetTags(build Build, github GitHub) []string {
	tags := []string{}
	if build.Tags != nil {
		tags = build.Tags
	}

	if !build.SetDefaultTags {
		return tags
	}

	switch github.Reference.Type {
	case GitRefHead:
		if github.Reference.Name == "master" {
			tags = append(tags, "latest")
		} else {
			tags = append(tags, github.Reference.Name)
		}
	case GitRefPullRequest:
		tags = append(tags, fmt.Sprintf("pr-%s", github.Reference.Name))
	case GitRefTag:
		tags = append(tags, github.Reference.Name)
	}

	return tags
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
