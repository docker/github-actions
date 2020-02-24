package options

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetBuildOptions(t *testing.T) {
	_ = os.Setenv("INPUT_PATH", "path")
	_ = os.Setenv("INPUT_DOCKERFILE", "dockerfile")
	_ = os.Setenv("INPUT_SERVER", "server")
	_ = os.Setenv("INPUT_REPOSITORY", "repository")
	_ = os.Setenv("INPUT_BUILD_ARGS", "buildarg1=b1,buildarg2=b2")
	_ = os.Setenv("INPUT_LABELS", "label1=l1,label2=l2")
	_ = os.Setenv("INPUT_AUTO_TAG", "false")
	_ = os.Setenv("INPUT_SET_DEFAULT_LABELS", "false")
	_ = os.Setenv("INPUT_TARGET", "target")
	_ = os.Setenv("INPUT_ALWAYS_PULL", "true")
	_ = os.Setenv("INPUT_TAGS", "tag1,tag2")

	o, err := GetBuildOptions()

	assert.NilError(t, err)
	assert.DeepEqual(t, Build{
		Path:             "path",
		Dockerfile:       "dockerfile",
		Server:           "server",
		SetDefaultLabels: false,
		Target:           "target",
		AlwaysPull:       true,
		BuildArgs:        []string{"buildarg1=b1", "buildarg2=b2"},
		Labels:           []string{"label1=l1", "label2=l2"},
	}, o)
}

func TestGetLabels(t *testing.T) {
	testCases := []struct {
		name       string
		setDefault bool
		labels     []string
		github     GitHub
		expected   []string
	}{
		{
			name:     "no-defaults",
			labels:   []string{"label1", "label2"},
			expected: []string{"label1", "label2"},
		},
		{
			name:       "with-defaults",
			labels:     []string{"label1", "label2"},
			setDefault: true,
			github: GitHub{
				Actor: "actor",
				Sha:   "sha",
			},
			expected: []string{
				"label1",
				"label2",
				"com.docker.github-actions-actor=actor",
				"com.docker.github-actions-sha=sha",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			labels := GetLabels(
				Build{
					SetDefaultLabels: tc.setDefault,
					Labels:           tc.labels,
				},
				tc.github,
			)
			assert.DeepEqual(t, tc.expected, labels)
		})
	}
}
