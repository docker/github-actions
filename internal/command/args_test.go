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
