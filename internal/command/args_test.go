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
	args := LoginArgs(o)

	assert.DeepEqual(t, expected, args)

	o.Server = "server"
	args = LoginArgs(o)
	expected = append(expected, "server")
	assert.DeepEqual(t, expected, args)
}

func TestBuildArgs(t *testing.T) {
	testCases := []struct {
		name     string
		build    options.Build
		github   options.GitHub
		expected []string
	}{
		{
			name:     "basic",
			expected: []string{"build", "."},
		},
		{
			name:     "with-dockerfile",
			build:    options.Build{Dockerfile: "dockerfile"},
			expected: []string{"build", "dockerfile"},
		},
		{
			name: "with-static-tags",
			build: options.Build{
				Repository: "repository",
				Tags:       []string{"tag1", "tag2"},
			},
			expected: []string{"build", "-t", "repository:tag1", "-t", "repository:tag2", "."},
		},
		{
			name: "with-static-tags-and-server",
			build: options.Build{
				Server:     "server",
				Repository: "repository",
				Tags:       []string{"tag1", "tag2"},
			},
			expected: []string{"build", "-t", "server/repository:tag1", "-t", "server/repository:tag2", "."},
		},
		{
			name: "with-default-tags",
			build: options.Build{
				SetDefaultTags: true,
				Repository:     "repository",
				Tags:           []string{"tag1"},
			},
			github: options.GitHub{
				Reference: options.GitReference{
					Type: options.GitRefHead,
					Name: "branch",
				},
			},
			expected: []string{"build", "-t", "repository:tag1", "-t", "repository:branch", "."},
		},
		{
			name: "with-static-labels",
			build: options.Build{
				Labels: []string{"label1", "label2"},
			},
			expected: []string{"build", "--label", "label1", "--label", "label2", "."},
		},
		{
			name: "with-default-labels",
			build: options.Build{
				SetDefaultLabels: true,
				Labels:           []string{"label1"},
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
				Target: "target",
			},
			expected: []string{"build", "--target", "target", "."},
		},
		{
			name: "with-always-pull",
			build: options.Build{
				AlwaysPull: true,
			},
			expected: []string{"build", "--pull", "."},
		},
		{
			name: "with-build-args",
			build: options.Build{
				BuildArgs: []string{"build-arg-1", "build-arg-2"},
			},
			expected: []string{"build", "--build-arg", "build-arg-1", "--build-arg", "build-arg-2", "."},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			args := BuildArgs(tc.build, tc.github)
			assert.DeepEqual(t, tc.expected, args)
		})
	}
}
