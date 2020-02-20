package options

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetGitHubOptions(t *testing.T) {
	_ = os.Setenv("GITHUB_ACTIONS", "true")
	_ = os.Setenv("GITHUB_WORKFLOW", "workflow")
	_ = os.Setenv("GITHUB_RUN_ID", "run-id")
	_ = os.Setenv("GITHUB_RUN_NUMBER", "run-number")
	_ = os.Setenv("GITHUB_ACTION", "action")
	_ = os.Setenv("GITHUB_ACTOR", "actor")
	_ = os.Setenv("GITHUB_REPOSITORY", "repository")
	_ = os.Setenv("GITHUB_EVENT_NAME", "event-name")
	_ = os.Setenv("GITHUB_SHA", "sha")
	_ = os.Setenv("GITHUB_REF", "ref")

	github, err := GetGitHubOptions()
	assert.NilError(t, err)
	assert.Equal(t, true, github.RunInActions)
	assert.Equal(t, "workflow", github.Workflow)
	assert.Equal(t, "run-id", github.RunID)
	assert.Equal(t, "run-number", github.RunNumber)
	assert.Equal(t, "action", github.Action)
	assert.Equal(t, "actor", github.Actor)
	assert.Equal(t, "repository", github.Repository)
	assert.Equal(t, "event-name", github.EventName)
	assert.Equal(t, "sha", github.Sha)
	assert.Equal(t, "ref", github.Ref)
}

func TestGetGitHubOptionsNotInActions(t *testing.T) {
	_ = os.Unsetenv("GITHUB_ACTIONS")
	github, err := GetGitHubOptions()
	assert.NilError(t, err)
	assert.Equal(t, false, github.RunInActions)
}
