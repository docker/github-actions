package e2e

import (
	"os/exec"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPush(t *testing.T) {
	err := setupLocalRegistry()
	assert.NilError(t, err)
	defer removeLocalRegistry()

	err = loginLocalRegistry()
	assert.NilError(t, err)

	tags := []string{"localhost:5000/my-repository:push-test"}

	err = removeImages(tags)
	assert.NilError(t, err)

	err = runActionsCommand("build", "testdata/push_tests/push.env")
	assert.NilError(t, err)

	err = runActionsCommand("push", "testdata/push_tests/push.env")
	assert.NilError(t, err)

	err = removeImages(tags)
	assert.NilError(t, err)

	err = exec.Command("docker", "pull", tags[0]).Run()
	defer removeImages(tags)
	assert.NilError(t, err)

	result, err := inspectImage(tags[0])
	assert.NilError(t, err)
	assert.DeepEqual(t,
		inspectResult{
			RepoTags: tags,
			Config: inspectResultConfig{
				Labels: map[string]string{
					"a": "a1",
					"b": "b1",
				},
			},
		}, result)
}
