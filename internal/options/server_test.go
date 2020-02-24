package options

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetServer(t *testing.T) {
	assert.Equal(t, "", GetServer())

	defer os.Unsetenv("INPUT_SERVER")
	_ = os.Setenv("INPUT_SERVER", "server")
	assert.Equal(t, "server", GetServer())
}
