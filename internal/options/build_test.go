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
	_ = os.Setenv("INPUT_SET_DEFAULT_TAGS", "false")
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
		Repository:       "repository",
		SetDefaultTags:   false,
		SetDefaultLabels: false,
		Target:           "target",
		AlwaysPull:       true,
		BuildArgs:        []string{"buildarg1=b1", "buildarg2=b2"},
		Labels:           []string{"label1=l1", "label2=l2"},
		Tags:             []string{"tag1", "tag2"},
	}, o)
}

func TestGetTags(t *testing.T) {
	testCases := []struct {
		name       string
		setDefault bool
		tags       []string
		ref        GitReference
		expected   []string
	}{
		{
			name:     "no-defaults",
			tags:     []string{"tag1", "tag2"},
			expected: []string{"tag1", "tag2"},
			ref:      GitReference{GitRefHead, "master"},
		},
		{
			name:       "unknown-ref-type",
			tags:       []string{"tag1", "tag2"},
			expected:   []string{"tag1", "tag2"},
			setDefault: true,
			ref:        GitReference{GitRefUnknown, "master"},
		},
		{
			name:       "master-branch",
			setDefault: true,
			tags:       []string{"tag1", "tag2"},
			expected:   []string{"tag1", "tag2", "latest"},
			ref:        GitReference{GitRefHead, "master"},
		},
		{
			name:       "different-branch",
			setDefault: true,
			tags:       []string{"tag1", "tag2"},
			expected:   []string{"tag1", "tag2", "branch-name"},
			ref:        GitReference{GitRefHead, "branch-name"},
		},
		{
			name:       "pull-request",
			setDefault: true,
			tags:       []string{"tag1", "tag2"},
			expected:   []string{"tag1", "tag2", "pr-name"},
			ref:        GitReference{GitRefPullRequest, "name"},
		},
		{
			name:       "tag",
			setDefault: true,
			tags:       []string{"tag1", "tag2"},
			expected:   []string{"tag1", "tag2", "v1.0"},
			ref:        GitReference{GitRefTag, "v1.0"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tags := GetTags(
				Build{
					SetDefaultTags: tc.setDefault,
					Tags:           tc.tags,
				},
				GitHub{Reference: tc.ref},
			)
			assert.DeepEqual(t, tc.expected, tags)
		})
	}
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
