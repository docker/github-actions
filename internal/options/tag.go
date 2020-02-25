package options

import (
	"fmt"
	"os"
	"strings"
)

func autoTag() bool {
	return os.Getenv("INPUT_AUTO_TAG") == "true"
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
	tag = strings.Trim(tag, " ")
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
	if autoTag() {
		switch github.Reference.Type {
		case GitRefHead:
			if github.Reference.Name == "master" {
				tags = append(tags, toFullTag(server, repo, "latest"))
			} else {
				tags = append(tags, toFullTag(server, repo, github.Reference.Name))
			}
			tags = appendShortGitShaTag(tags, github, server, repo)
		case GitRefPullRequest:
			tags = append(tags, toFullTag(server, repo, fmt.Sprintf("pr-%s", github.Reference.Name)))
			tags = appendShortGitShaTag(tags, github, server, repo)
		case GitRefTag:
			tags = append(tags, toFullTag(server, repo, github.Reference.Name))
		}
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
