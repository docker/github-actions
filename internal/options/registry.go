package options

import "os"

// GetRegistry gets the registry server from the github actions environment variables
func GetRegistry() string {
	return os.Getenv("INPUT_REGISTRY")
}
