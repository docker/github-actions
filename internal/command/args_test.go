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

	args = LoginArgs(o, "server")
	expected = append(expected, "server")
	assert.DeepEqual(t, expected, args)
}

func TestBuildArgs(t *testing.T) {
	testCases := []struct {
		name     string
		build    options.Build
		github   options.GitHub
		tags     []string
		expected []string
	}{
		{
			name:     "basic",
			build:    options.Build{Path: "path"},
			expected: []string{"build", "path"},
		},
		{
			name:     "with-dockerfile",
			build:    options.Build{Path: ".", Dockerfile: "dockerfile"},
			expected: []string{"build", "--file", "dockerfile", "."},
		},
		{
			name:     "with-tags",
			build:    options.Build{Path: "."},
			tags:     []string{"tag1", "tag2"},
			expected: []string{"build", "-t", "tag1", "-t", "tag2", "."},
		},
		{
			name: "with-static-labels",
			build: options.Build{
				Path:   ".",
				Labels: []string{"label1", "label2"},
			},
			expected: []string{"build", "--label", "label1", "--label", "label2", "."},
		},
		{
			name: "with-git-labels",
			build: options.Build{
				Path:         ".",
				AddGitLabels: true,
				Labels:       []string{"label1"},
			},
			github: options.GitHub{
				Actor: "actor",
				Sha:   "sha",
			},
			expected: []string{"build", "--label", "label1", "--label", "com.docker.github-actions-actor=actor", "--label", "com.docker.github-actions-sha=sha", "."},
		},
		{
			name: "with-target",
			build: options.Build{
				Path:   ".",
				Target: "target",
			},
			expected: []string{"build", "--target", "target", "."},
		},
		{
			name: "with-always-pull",
			build: options.Build{
				Path:       ".",
				AlwaysPull: true,
			},
			expected: []string{"build", "--pull", "."},
		},
		{
			name: "with-build-args",
			build: options.Build{
				Path:      ".",
				BuildArgs: []string{"build-arg-1", "build-arg-2"},
			},
			expected: []string{"build", "--build-arg", "build-arg-1", "--build-arg", "build-arg-2", "."},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			args := BuildArgs(tc.build, tc.github, tc.tags)
			assert.DeepEqual(t, tc.expected, args)
		})
	}
}

func TestPushArgs(t *testing.T) {
	args := PushArgs("tag1")
	assert.DeepEqual(t, []string{"push", "tag1"}, args)
}
