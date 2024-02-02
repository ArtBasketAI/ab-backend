package config

import (
	"errors"

	"github.com/kelseyhightower/envconfig"

)

type Env struct {
	// LogLevel is INFO or DEBUG. Default is "INFO".
	LogLevel string `envconfig:"LOG_LEVEL" default:"INFO"`
	// Env is used to set GIN log level
	Env string `envconfig:"ENV" default:"development"`
	// ServiceName used in setting parameters as ID for various monitoring tools
	ServiceName string `envconfig:"SERVICE_NAME" default:"ab-backend"`
	// Version is a version of the application binary.
	Version string `envconfig:"VERSION"`
}

// ReadFromEnv reads configuration from environmental variables
// defined by Env struct.
func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Join(err, errors.New("failed to process envconfig"))
	}

	return &env, nil
}

func GetPort() string {
	return "8080"
}
