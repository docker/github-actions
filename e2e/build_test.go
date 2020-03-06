package e2e

import (
	"os/exec"
	"testing"
	"time"

	"github.com/hashicorp/go-multierror"
	"gotest.tools/v3/assert"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		name           string
		envFile        string
		expectedTags   []string
		expectedLabels map[string]string
	}{
		{
			name:    "static-tags",
			envFile: "testdata/build_tests/static_tags.env",
			expectedTags: []string{
				"localhost:5000/my-repository:v1-static-tags",
				"localhost:5000/my-repository:v1.1-static-tags",
			},
		},
		{
			name:    "static-labels",
			envFile: "testdata/build_tests/static_labels.env",
			expectedTags: []string{
				"localhost:5000/my-repository:static-labels",
			},
			expectedLabels: map[string]string{
				"a": "a1",
				"b": "b1",
			},
		},
		{
			name:    "auto-tags-master",
			envFile: "testdata/build_tests/tag_master.env",
			expectedTags: []string{
				"localhost:5000/my-repository:auto-tags-master",
				"localhost:5000/my-repository:latest",
			},
		},
		{
			name:    "auto-tags-branch",
			envFile: "testdata/build_tests/tag_branch.env",
			expectedTags: []string{
				"localhost:5000/my-repository:auto-tags-branch",
				"localhost:5000/my-repository:branch",
			},
		},
		{
			name:    "auto-tags-pr",
			envFile: "testdata/build_tests/tag_pr.env",
			expectedTags: []string{
				"localhost:5000/my-repository:pr-pr1",
			},
		},
		{
			name:    "auto-tags-tag",
			envFile: "testdata/build_tests/tag_tag.env",
			expectedTags: []string{
				"localhost:5000/my-repository:tag1",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := runActionsCommand("build", tc.envFile)
			assert.NilError(t, err)
			defer removeImages(tc.expectedTags)

			for _, tag := range tc.expectedTags {
				inspect, err := inspectImage(tag)
				assert.NilError(t, err)
				assert.DeepEqual(t, tc.expectedTags, inspect.RepoTags)
				assert.DeepEqual(t, tc.expectedLabels, inspect.Config.Labels)
			}
		})
	}
}

func TestBuildWithGitLabels(t *testing.T) {
	tags := []string{"localhost:5000/my-repository:auto-labels"}
	err := runActionsCommand("build", "testdata/build_tests/auto_labels.env")
	assert.NilError(t, err)
	defer removeImages(tags)

	inspect, err := inspectImage(tags[0])
	assert.NilError(t, err)
	assert.DeepEqual(t, tags, inspect.RepoTags)

	assert.Equal(t, 4, len(inspect.Config.Labels))
	assert.Equal(t, "a1", inspect.Config.Labels["a"])
	assert.Equal(t, "https://github.com/git/repository", inspect.Config.Labels["org.opencontainers.image.source"])
	assert.Equal(t, "sha", inspect.Config.Labels["org.opencontainers.image.revision"])

	created, err := time.Parse(time.RFC3339, inspect.Config.Labels["org.opencontainers.image.created"])
	assert.NilError(t, err)
	assert.Assert(t, created.Before(time.Now()))
}

func removeImages(tags []string) error {
	var result error
	for _, tag := range tags {
		if err := exec.Command("docker", "rmi", "-f", tag).Run(); err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result
}
