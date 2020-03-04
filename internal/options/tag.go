package options

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tagWithRef() bool {
	b, err := strconv.ParseBool(os.Getenv("INPUT_TAG_WITH_REF"))
	return err == nil && b
}

func tagWithSha() bool {
	b, err := strconv.ParseBool(os.Getenv("INPUT_TAG_WITH_SHA"))
	return err == nil && b
}

func dockerRepo(github GitHub) string {
	if repo := os.Getenv("INPUT_REPOSITORY"); repo != "" {
		return repo
	}
	return strings.ToLower(github.Repository)
}

func staticTags() []string {
	if inputTags := os.Getenv("INPUT_TAGS"); inputTags != "" {
		return strings.Split(inputTags, ",")
	}
	return nil
}

func toFullTag(registry, repo, tag string) string {
	tag = strings.TrimSpace(tag)
	tag = strings.ReplaceAll(tag, "/", "-")
	if registry != "" {
		return fmt.Sprintf("%s/%s:%s", registry, repo, tag)
	}
	return fmt.Sprintf("%s:%s", repo, tag)
}

// GetTags gets a list of all tags for including automatic tags from github vars when enabled along with the registry and repository
func GetTags(registry string, github GitHub) []string {
	repo := dockerRepo(github)
	var tags []string
	for _, t := range staticTags() {
		tags = append(tags, toFullTag(registry, repo, t))
	}
	if tagWithRef() {
		switch github.Reference.Type {
		case GitRefHead:
			if github.Reference.Name == "master" {
				tags = append(tags, toFullTag(registry, repo, "latest"))
			} else {
				tags = appendGitRefTag(tags, registry, repo, github.Reference.Name)
			}
		case GitRefPullRequest:
			tags = appendGitRefTag(tags, registry, repo, fmt.Sprintf("pr-%s", github.Reference.Name))

		case GitRefTag:
			tags = appendGitRefTag(tags, registry, repo, github.Reference.Name)
		}
	}
	if tagWithSha() {
		tags = appendShortGitShaTag(tags, github, registry, repo)
	}
	return tags
}

func appendShortGitShaTag(tags []string, github GitHub, registry, repo string) []string {
	if len(github.Sha) >= 7 {
		tag := fmt.Sprintf("sha-%s", github.Sha[0:7])
		return append(tags, toFullTag(registry, repo, tag))
	}
	return tags
}

func appendGitRefTag(tags []string, registry, repo, refName string) []string {
	t := strings.ReplaceAll(refName, "/", "-")
	return append(tags, toFullTag(registry, repo, t))
}
