package options

import (
	"errors"

	"github.com/caarlos0/env/v6"
)

// Login contains the parsed login action environment variables
type Login struct {
	Username string `env:"INPUT_USERNAME"`
	Password string `env:"INPUT_PASSWORD"`
}

var errLoginVarValidation = errors.New("both username and password must be set to login")

// GetLoginOptions gets the login action environment variables
func GetLoginOptions() (Login, error) {
	var login Login
	if err := env.Parse(&login); err != nil {
		return login, err
	}

	if login.Username != "" && login.Password == "" ||
		login.Username == "" && login.Password != "" {
		return login, errLoginVarValidation
	}

	return login, nil
}
