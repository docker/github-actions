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

func readUint64Option(key string) (uint64, error) {
	o := os.Getenv(key)
	i, err := strconv.Atoi(o)
	if err != nil {
		return 0, err
	}
	return uint64(i), nil
}
