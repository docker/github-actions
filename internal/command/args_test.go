package command

import (
	"testing"

	"github.com/docker/github-actions/internal/options"
	"gotest.tools/v3/assert"
)

func TestLoginArgs(t *testing.T) {
	expected := []string{"login", "--username", "username", "--password", "password"}
	o := options.Login{
		Username: "username",
		Password: "password",
	}
	args := LoginArgs(o, "")

	assert.DeepEqual(t, expected, args)

	args = LoginArgs(o, "registry")
	expected = append(expected, "registry")
	assert.DeepEqual(t, expected, args)
}

func TestBuildArgs(t *testing.T) {
	testCases := []struct {
		name     string
		build    options.Build
		tags     []string
		expected []string
	}{
		{
			name:     "basic",
			build:    options.Build{Path: "path"},
			expected: []string{"build", "--progress", "plain", "path"},
		},
		{
			name:     "with-dockerfile",
			build:    options.Build{Path: ".", Dockerfile: "dockerfile"},
			expected: []string{"build", "--progress", "plain", "--file", "dockerfile", "."},
		},
		{
			name:     "with-tags",
			build:    options.Build{Path: "."},
			tags:     []string{"tag1", "tag2"},
			expected: []string{"build", "--progress", "plain", "-t", "tag1", "-t", "tag2", "."},
		},
		{
			name: "with-static-labels",
			build: options.Build{
				Path:   ".",
				Labels: []string{"label1", "label2"},
			},
			expected: []string{"build", "--progress", "plain", "--label", "label1", "--label", "label2", "."},
		},
		{
			name: "with-target",
			build: options.Build{
				Path:   ".",
				Target: "target",
			},
			expected: []string{"build", "--progress", "plain", "--target", "target", "."},
		},
		{
			name: "with-always-pull",
			build: options.Build{
				Path:       ".",
				AlwaysPull: true,
			},
			expected: []string{"build", "--progress", "plain", "--pull", "."},
		},
		{
			name: "with-build-args",
			build: options.Build{
				Path:      ".",
				BuildArgs: []string{"build-arg-1", "build-arg-2"},
			},
			expected: []string{"build", "--progress", "plain", "--build-arg", "build-arg-1", "--build-arg", "build-arg-2", "."},
		},
		{
			name: "with-cache-from",
			build: options.Build{
				Path:       ".",
				CacheFroms: []string{"foo/bar-1", "foo/bar-2"},
			},
			expected: []string{"build", "--progress", "plain", "--cache-from", "foo/bar-1", "--cache-from", "foo/bar-2", "."},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			args := BuildArgs(tc.build, options.GitHub{}, tc.tags)
			assert.DeepEqual(t, tc.expected, args)
		})
	}
}

func TestPushArgs(t *testing.T) {
	args := PushArgs("tag1")
	assert.DeepEqual(t, []string{"push", "tag1"}, args)
}
