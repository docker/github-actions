package options

import (
	"os"
	"strconv"
)

func readBoolOption(key string) (bool, error) {
	o := os.Getenv(key)
	if o == "" {
		return false, nil
	}
	return strconv.ParseBool(o)
}
