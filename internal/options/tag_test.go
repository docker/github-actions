package options

import (
	"fmt"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetTags(t *testing.T) {
	testCases := []struct {
		name       string
		tagWithRef bool
		tagWithSha bool
		tags       string
		ref        GitReference
		server     string
		expected   []string
		sha        string
	}{
		{
			name:     "no-standard-tags",
			tags:     "tag1,tag2",
			expected: []string{"my/repo:tag1", "my/repo:tag2"},
			ref:      GitReference{GitRefHead, "master"},
		},
		{
			name:     "with-server",
			tags:     "tag1,tag2",
			expected: []string{"server/my/repo:tag1", "server/my/repo:tag2"},
			server:   "server",
			ref:      GitReference{GitRefHead, "master"},
		},
		{
			name:       "unknown-ref-type",
			tags:       "tag1,tag2",
			expected:   []string{"my/repo:tag1", "my/repo:tag2"},
			tagWithRef: true,
			ref:        GitReference{GitRefUnknown, "master"},
		},
		{
			name:       "master-branch",
			tagWithRef: true,
			tags:       "tag1,tag2",
			expected:   []string{"my/repo:tag1", "my/repo:tag2", "my/repo:latest"},
			ref:        GitReference{GitRefHead, "master"},
		},
		{
			name:       "different-branch",
			tagWithRef: true,
			tags:       "tag1,tag2",
			expected:   []string{"my/repo:tag1", "my/repo:tag2", "my/repo:branch-name"},
			ref:        GitReference{GitRefHead, "branch-name"},
		},
		{
			name:       "pull-request",
			tagWithRef: true,
			tags:       "tag1,tag2",
			expected:   []string{"my/repo:tag1", "my/repo:tag2", "my/repo:pr-name"},
			ref:        GitReference{GitRefPullRequest, "name"},
		},
		{
			name:       "tag",
			tagWithRef: true,
			tags:       "tag1,tag2",
			expected:   []string{"my/repo:tag1", "my/repo:tag2", "my/repo:v1.0"},
			ref:        GitReference{GitRefTag, "v1.0"},
		},
		{
			name:       "master-branch-with-sha",
			tagWithRef: true,
			tags:       "tag1,tag2",
			expected:   []string{"my/repo:tag1", "my/repo:tag2", "my/repo:latest"},
			ref:        GitReference{GitRefHead, "master"},
			sha:        "1234567890",
		},
		{
			name:       "pull-request-with-sha",
			tagWithRef: true,
			tags:       "tag1,tag2",
			expected:   []string{"my/repo:tag1", "my/repo:tag2", "my/repo:pr-name"},
			ref:        GitReference{GitRefPullRequest, "name"},
			sha:        "1234567890",
		},
		{
			name:       "tag-with-sha",
			tagWithSha: true,
			tags:       "tag1,tag2",
			expected:   []string{"my/repo:tag1", "my/repo:tag2", "my/repo:sha-1234567"},
			sha:        "1234567890",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			defer os.Unsetenv("INPUT_TAGS")
			defer os.Unsetenv("INPUT_REPOSITORY")
			defer os.Unsetenv("INPUT_TAG_WITH_REF")
			defer os.Unsetenv("INPUT_TAG_WITH_SHA")
			_ = os.Setenv("INPUT_TAGS", tc.tags)
			_ = os.Setenv("INPUT_REPOSITORY", "my/repo")
			_ = os.Setenv("INPUT_TAG_WITH_REF", fmt.Sprint(tc.tagWithRef))
			_ = os.Setenv("INPUT_TAG_WITH_SHA", fmt.Sprint(tc.tagWithSha))

			tags := GetTags(
				tc.server,
				GitHub{Reference: tc.ref, Sha: tc.sha},
			)
			assert.DeepEqual(t, tc.expected, tags)
		})
	}
}

func TestGetTagsWithGitHubRepo(t *testing.T) {
	defer os.Unsetenv("INPUT_TAGS")
	_ = os.Setenv("INPUT_TAGS", "tag1")

	github := GitHub{Repository: "My/Repo"}
	tags := GetTags("", github)
	assert.DeepEqual(t, []string{"my/repo:tag1"}, tags)
}
