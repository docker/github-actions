package e2e

import (
	"os/exec"
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

func testBuildPush(t *testing.T, envFile string, tags []string, labels map[string]string) {
	err := removeImages(tags)
	assert.NilError(t, err)

	err = runActionsCommand("build-push", envFile)
	assert.NilError(t, err)

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

func TestBuildPush(t *testing.T) {
	err := setupLocalRegistry()
	assert.NilError(t, err)
	defer removeLocalRegistry()

	err = ensureLocalRegistryAlive()
	assert.NilError(t, err)

	// Build and push base image
	baseTags := []string{
		"localhost:5000/org/base:build-push-tag1",
		"localhost:5000/org/base:build-push-test",
	}
	defer removeImages(baseTags)
	testBuildPush(
		t,
		"testdata/build_push_tests/build_push.env",
		baseTags,
		map[string]string{
			"a": "a1",
		},
	)

	err = logoutLocalRegistry()
	assert.NilError(t, err)

	// Build and push image using base image from local registry
	testBuildPush(
		t,
		"testdata/build_push_tests/build_push_from_registry.env",
		[]string{
			"localhost:5000/org/repo:build-push-reg-tag1",
			"localhost:5000/org/repo:build-push-reg-test",
		},
		map[string]string{
			"a": "a1",
			"b": "b1",
		},
	)
}

func assertBuildPushImages(t *testing.T, image string, expectedTags []string, expectedLabels map[string]string) {
	inspect, err := inspectImage(image)
	assert.NilError(t, err)

	repoTags := inspect.RepoTags
	sort.Strings(repoTags)

	assert.DeepEqual(t, expectedTags, repoTags)
	assert.DeepEqual(t, expectedLabels, inspect.Config.Labels)
}

func ensureLocalRegistryAlive() error {
	if err := loginLocalRegistry(); err != nil {
		return err
	}

	return logoutLocalRegistry()
}

func logoutLocalRegistry() error {
	return exec.Command("docker", "logout", "localhost:5000").Run()
}
