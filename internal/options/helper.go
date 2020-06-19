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

// OptionalBool enum
type OptionalBool int

const (
	notPresent OptionalBool = iota
	presentAndFalse
	presentAndTrue
)

func (d OptionalBool) String() string {
	return [...]string{"notPresent", "presentAndFalse", "presentAndTrue"}[d]
}
