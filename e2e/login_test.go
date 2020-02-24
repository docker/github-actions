package e2e

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestLogin(t *testing.T) {
	err := setupLocalRegistry()
	assert.NilError(t, err)
	defer removeLocalRegistry()

	err = runActionsCommand("login", "testdata/login_test.env")
	assert.NilError(t, err)
}
