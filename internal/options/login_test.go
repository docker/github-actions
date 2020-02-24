package options

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetLoginOptions(t *testing.T) {
	_ = os.Setenv("INPUT_USERNAME", "username")
	_ = os.Setenv("INPUT_PASSWORD", "password")

	o, err := GetLoginOptions()
	assert.NilError(t, err)
	assert.Equal(t, "username", o.Username)
	assert.Equal(t, "password", o.Password)
}
