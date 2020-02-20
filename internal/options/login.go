package options

import "github.com/caarlos0/env/v6"

// Login contains the parsed login action environment variables
type Login struct {
	Username string `env:"INPUT_USERNAME"`
	Password string `env:"INPUT_PASSWORD"`
	Server   string `env:"INPUT_SERVER"`
}

// GetLoginOptions gets the login action environment variables
func GetLoginOptions() (Login, error) {
	var login Login
	err := env.Parse(&login)
	return login, err
}
