package options

import (
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
)

const opencontainersLabelPrefix = "org.opencontainers.image"

// Build contains the parsed build action environment variables
type Build struct {
	Path         string `env:"INPUT_PATH"`
	Dockerfile   string `env:"INPUT_DOCKERFILE"`
	AddGitLabels bool   `env:"INPUT_ADD_GIT_LABELS"`
	Target       string `env:"INPUT_TARGET"`
	AlwaysPull   bool   `env:"INPUT_ALWAYS_PULL"`
	CacheFroms   []string
	BuildArgs    []string
	Labels       []string
}

// GetBuildOptions gets the login action environment variables
func GetBuildOptions() (Build, error) {
	var build Build
	if err := env.Parse(&build); err != nil {
		return build, err
	}

	if cacheFroms := os.Getenv("INPUT_CACHE_FROMS"); cacheFroms != "" {
		build.CacheFroms = strings.Split(cacheFroms, ",")
	}

	if buildArgs := os.Getenv("INPUT_BUILD_ARGS"); buildArgs != "" {
		build.BuildArgs = strings.Split(buildArgs, ",")
	}

	if labels := os.Getenv("INPUT_LABELS"); labels != "" {
		build.Labels = strings.Split(labels, ",")
	}

	return build, nil
}

// GetLabels gets a list of all labels to build the image with including automatic labels created from github vars when AddGitLabels is true
func GetLabels(build Build, github GitHub) []string {
	labels := []string{}
	if build.Labels != nil {
		labels = build.Labels
	}

	if !build.AddGitLabels {
		return labels
	}

	if github.Repository != "" {
		labels = append(labels, opencontainersLabelPrefix+".source=https://github.com/"+github.Repository)
	}

	if github.Sha != "" {
		labels = append(labels, opencontainersLabelPrefix+".revision="+github.Sha)

	}

	createdTime := time.Now().UTC().Format(time.RFC3339)
	labels = append(labels, opencontainersLabelPrefix+".created="+createdTime)

	return labels
}
