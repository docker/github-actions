package options

import (
	"errors"
)

var errPushParse = errors.New("push input must be a valid boolean value")

// ShouldPush returns true if the user has signalled a docker push should be performed
func ShouldPush() (bool, error) {
	b, err := readBoolOption("INPUT_PUSH")
	if err != nil {
		return false, errPushParse
	}
	return b, nil
}

// MaxPushRetries returns the maximum number of push retries that should be performed or 0
func MaxPushRetries() uint64 {
	b, err := readUint64Option("INPUT_MAX_PUSH_RETRIES")
	if err != nil {
		return 0
	}
	return b
}
