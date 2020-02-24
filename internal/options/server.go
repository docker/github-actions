package options

import "os"

// GetServer gets the server from the github actions environment variables
func GetServer() string {
	return os.Getenv("INPUT_SERVER")
}
