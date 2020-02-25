package e2e

import (
	"os/exec"
	"testing"

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
			name:    "auto-labels",
			envFile: "testdata/build_tests/auto_labels.env",
			expectedTags: []string{
				"localhost:5000/my-repository:auto-labels",
			},
			expectedLabels: map[string]string{
				"a":                               "a1",
				"com.docker.github-actions-actor": "actor",
				"com.docker.github-actions-sha":   "sha",
			},
		},
		{
			name:    "auto-tags-master",
			envFile: "testdata/build_tests/auto_tags_master.env",
			expectedTags: []string{
				"localhost:5000/my-repository:auto-tags-master",
				"localhost:5000/my-repository:latest",
			},
		},
		{
			name:    "auto-tags-branch",
			envFile: "testdata/build_tests/auto_tags_branch.env",
			expectedTags: []string{
				"localhost:5000/my-repository:auto-tags-branch",
				"localhost:5000/my-repository:branch",
			},
		},
		{
			name:    "auto-tags-pr",
			envFile: "testdata/build_tests/auto_tags_pr.env",
			expectedTags: []string{
				"localhost:5000/my-repository:pr-pr1",
			},
		},
		{
			name:    "auto-tags-tag",
			envFile: "testdata/build_tests/auto_tags_tag.env",
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

func removeImages(tags []string) error {
	for _, tag := range tags {
		if err := exec.Command("docker", "rmi", "-f", tag).Run(); err != nil {
			return err
		}
	}
	return nil
}
