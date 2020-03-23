package e2e

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"
	"k8s.io/apimachinery/pkg/util/wait"
)

func TestLogin(t *testing.T) {
	err := setupLocalRegistry()
	assert.NilError(t, err)
	defer removeLocalRegistry()

	err = loginLocalRegistry()
	assert.NilError(t, err)
}

func loginLocalRegistry() error {
	// Polls as registry takes a moment to start up
	return wait.Poll(2*time.Second, 30*time.Second, func() (bool, error) {
		err := runActionsCommand("login", "testdata/login_test.env")
		return err == nil, err
	})
}
