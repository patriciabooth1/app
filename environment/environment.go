// Package environment provides access to the environment
// variables required by the server.
package environment

import (
	"github.com/kelseyhightower/envconfig"
)

// Environment represents a set of validated environment variables.
type Environment struct {
	// Port is the port to the run server on.
	Port int `required:"true"`
	// LogLevel is the level of logging required.
	LogLevel string `split_words:"true"`
}

// Get returns the current environment.
func Get() (*Environment, error) {
	var e Environment
	err := envconfig.Process("", &e)
	return &e, err
}
