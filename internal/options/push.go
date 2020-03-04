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
