package options

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestShouldPush(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty",
			input:    "",
			expected: false,
		},
		{
			name:     "true",
			input:    "true",
			expected: true,
		},
		{
			name:     "false",
			input:    "false",
			expected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_ = os.Setenv("INPUT_PUSH", tc.input)
			defer os.Unsetenv("INPUT_PUSH")
			assert.Equal(t, tc.expected, ShouldPush())
		})
	}
}
