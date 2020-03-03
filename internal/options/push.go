package options

import (
	"os"
	"strconv"
)

// ShouldPush returns true if the user has signalled a docker push should be performed
func ShouldPush() bool {
	b, err := strconv.ParseBool(os.Getenv("INPUT_PUSH"))
	return err == nil && b
}
