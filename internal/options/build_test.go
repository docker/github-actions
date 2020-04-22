package options

import (
	"os"
	"strings"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestGetBuildOptions(t *testing.T) {
	_ = os.Setenv("INPUT_PATH", "path")
	_ = os.Setenv("INPUT_DOCKERFILE", "dockerfile")
	_ = os.Setenv("INPUT_CACHE_FROMS", "foo/bar-1,foo/bar-2")
	_ = os.Setenv("INPUT_REPOSITORY", "repository")
	_ = os.Setenv("INPUT_BUILD_ARGS", "buildarg1=b1,buildarg2=b2")
	_ = os.Setenv("INPUT_LABELS", "label1=l1,label2=l2")
	_ = os.Setenv("INPUT_ADD_GIT_LABELS", "false")
	_ = os.Setenv("INPUT_TARGET", "target")
	_ = os.Setenv("INPUT_ALWAYS_PULL", "true")

	o, err := GetBuildOptions()

	assert.NilError(t, err)
	assert.DeepEqual(t, Build{
		Path:       "path",
		Dockerfile: "dockerfile",
		Target:     "target",
		AlwaysPull: true,
		BuildArgs:  []string{"buildarg1=b1", "buildarg2=b2"},
		Labels:     []string{"label1=l1", "label2=l2"},
		CacheFroms: []string{"foo/bar-1", "foo/bar-2"},
	}, o)
}

func TestGetLabels(t *testing.T) {
	expected := []string{"label1", "label2"}
	labels := GetLabels(Build{Labels: expected}, GitHub{})
	assert.DeepEqual(t, expected, labels)
}

func TestGetLabelsWithGit(t *testing.T) {
	labels := GetLabels(
		Build{
			Labels:       []string{"label1", "label2"},
			AddGitLabels: true,
		},
		GitHub{
			Repository: "myrepository",
			Sha:        "mysha",
		})

	assert.Equal(t, 5, len(labels))
	assert.Equal(t, "label1", labels[0])
	assert.Equal(t, "label2", labels[1])
	assert.Equal(t, "org.opencontainers.image.source=https://github.com/myrepository", labels[2])
	assert.Equal(t, "org.opencontainers.image.revision=mysha", labels[3])

	timeLabel := strings.SplitN(labels[4], "=", 2)
	assert.Equal(t, "org.opencontainers.image.created", timeLabel[0])
	labelTime, err := time.Parse(time.RFC3339, timeLabel[1])
	assert.NilError(t, err)
	assert.Assert(t, time.Now().After(labelTime))
}
