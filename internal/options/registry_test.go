package options

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetRegistry(t *testing.T) {
	_ = os.Unsetenv("INPUT_REGISTRY")
	assert.Equal(t, "", GetRegistry())

	defer os.Unsetenv("INPUT_REGISTRY")
	_ = os.Setenv("INPUT_REGISTRY", "registry")
	assert.Equal(t, "registry", GetRegistry())
}
