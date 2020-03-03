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
	return []string{}
}

func toFullTag(server, repo, tag string) string {
	tag = strings.TrimSpace(tag)
	tag = strings.ReplaceAll(tag, "/", "-")
	if server != "" {
		return fmt.Sprintf("%s/%s:%s", server, repo, tag)
	}
	return fmt.Sprintf("%s:%s", repo, tag)
}

// GetTags gets a list of all tags for including automatic tags created from github vars when AutTag is true along with the server and repository
func GetTags(server string, github GitHub) []string {
	repo := dockerRepo(github)
	tags := []string{}
	for _, t := range staticTags() {
		tags = append(tags, toFullTag(server, repo, t))
	}
	if tagWithRef() {
		switch github.Reference.Type {
		case GitRefHead:
			if github.Reference.Name == "master" {
				tags = append(tags, toFullTag(server, repo, "latest"))
			} else {
				tags = appendGitRefTag(tags, server, repo, github.Reference.Name)
			}
		case GitRefPullRequest:
			tags = appendGitRefTag(tags, server, repo, fmt.Sprintf("pr-%s", github.Reference.Name))

		case GitRefTag:
			tags = appendGitRefTag(tags, server, repo, github.Reference.Name)
		}
	}
	if tagWithSha() {
		tags = appendShortGitShaTag(tags, github, server, repo)
	}
	return tags
}

func appendShortGitShaTag(tags []string, github GitHub, server, repo string) []string {
	if len(github.Sha) >= 7 {
		tag := fmt.Sprintf("sha-%s", github.Sha[0:7])
		return append(tags, toFullTag(server, repo, tag))
	}
	return tags
}

func appendGitRefTag(tags []string, server, repo, refName string) []string {
	t := strings.ReplaceAll(refName, "/", "-")
	return append(tags, toFullTag(server, repo, t))
}
