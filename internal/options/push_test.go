package options

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestShouldPush(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    bool
		expectedErr error
	}{
		{
			name:        "invalid",
			input:       "invalid",
			expectedErr: errPushParse,
		},
		{
			name:     "true",
			input:    "true",
			expected: true,
		},
		{
			name:  "false",
			input: "false",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_ = os.Setenv("INPUT_PUSH", tc.input)
			defer os.Unsetenv("INPUT_PUSH")
			should, err := ShouldPush()
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expected, should)
		})
	}
}
