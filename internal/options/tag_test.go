package options

import (
	"fmt"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetTags(t *testing.T) {
	testCases := []struct {
		name          string
		tagWithRef    bool
		tagWithSha    bool
		tagWithLatest OptionalBool
		tags          string
		ref           GitReference
		registry      string
		expected      []string
		sha           string
	}{
		{
			name:     "no-standard-tags",
			tags:     "tag1,tag2",
			expected: []string{"my/repo:tag1", "my/repo:tag2"},
			ref:      GitReference{GitRefHead, "master"},
		},
		{
			name:     "with-registry",
			tags:     "tag1,tag2",
			expected: []string{"registry/my/repo:tag1", "registry/my/repo:tag2"},
			registry: "registry",
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
			name:          "latest-true",
			tags:          "tag1,tag2",
			tagWithLatest: presentAndTrue,
			expected:      []string{"my/repo:tag1", "my/repo:tag2", "my/repo:latest"},
			ref:           GitReference{GitRefHead, "master"},
		},
		{
			name:          "latest-false",
			tags:          "tag1,tag2",
			tagWithLatest: presentAndFalse,
			expected:      []string{"my/repo:tag1", "my/repo:tag2"},
			ref:           GitReference{GitRefHead, "master"},
		},
		{
			name:          "master-branch-latest-true",
			tagWithRef:    true,
			tags:          "tag1,tag2",
			tagWithLatest: presentAndTrue,
			expected:      []string{"my/repo:tag1", "my/repo:tag2", "my/repo:latest", "my/repo:master"},
			ref:           GitReference{GitRefHead, "master"},
		},
		{
			name:          "master-branch-latest-false",
			tagWithRef:    true,
			tags:          "tag1,tag2",
			tagWithLatest: presentAndFalse,
			expected:      []string{"my/repo:tag1", "my/repo:tag2", "my/repo:master"},
			ref:           GitReference{GitRefHead, "master"},
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
			defer os.Unsetenv("INPUT_TAG_WITH_LATEST")
			_ = os.Setenv("INPUT_TAGS", tc.tags)
			_ = os.Setenv("INPUT_REPOSITORY", "my/repo")
			_ = os.Setenv("INPUT_TAG_WITH_REF", fmt.Sprint(tc.tagWithRef))
			_ = os.Setenv("INPUT_TAG_WITH_SHA", fmt.Sprint(tc.tagWithSha))
			if tc.tagWithLatest == presentAndTrue || tc.tagWithLatest == presentAndFalse {
				if tc.tagWithLatest == presentAndTrue {
					_ = os.Setenv("INPUT_TAG_WITH_LATEST", fmt.Sprint(true))
				} else {
					_ = os.Setenv("INPUT_TAG_WITH_LATEST", fmt.Sprint(false))
				}
			}

			tags, err := GetTags(
				tc.registry,
				GitHub{Reference: tc.ref, Sha: tc.sha},
			)
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.expected, tags)
		})
	}
}

func TestGetTagsWithGitHubRepo(t *testing.T) {
	defer os.Unsetenv("INPUT_TAGS")
	_ = os.Setenv("INPUT_TAGS", "tag1")

	github := GitHub{Repository: "My/Repo"}
	tags, err := GetTags("", github)
	assert.NilError(t, err)
	assert.DeepEqual(t, []string{"my/repo:tag1"}, tags)
}
