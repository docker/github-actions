package e2e

import (
	"os/exec"
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

func TestBuildPush(t *testing.T) {
	tags := []string{
		"localhost:5000/my-repository:build-push-tag1",
		"localhost:5000/my-repository:build-push-test",
	}
	labels := map[string]string{
		"a": "a1",
	}
	err := removeImages(tags)
	assert.NilError(t, err)
	defer removeImages(tags)

	err = setupLocalRegistry()
	assert.NilError(t, err)
	defer removeLocalRegistry()

	err = loginLocalRegistry()
	assert.NilError(t, err)

	err = runActionsCommand("build-push", "testdata/build_push_tests/build_push.env")
	assert.NilError(t, err)

	for _, tag := range tags {
		assertBuildPushImages(t, tag, tags, labels)
	}

	err = removeImages(tags)
	assert.NilError(t, err)

	for _, tag := range tags {
		err = exec.Command("docker", "pull", tag).Run()
		assert.NilError(t, err)
	}

	for _, tag := range tags {
		assertBuildPushImages(t, tag, tags, labels)
	}
}

func assertBuildPushImages(t *testing.T, image string, expectedTags []string, expectedLabels map[string]string) {
	inspect, err := inspectImage(image)
	assert.NilError(t, err)

	repoTags := inspect.RepoTags
	sort.Strings(repoTags)

	assert.DeepEqual(t, expectedTags, repoTags)
	assert.DeepEqual(t, expectedLabels, inspect.Config.Labels)
}
