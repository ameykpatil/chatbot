package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Specification structured configuration variables
type Specification struct {
	HTTP struct {
		BaseURL string `envconfig:"HTTP_BASE_URL"`
		Port    int    `envconfig:"HTTP_PORT" default:"8000"`
	}
	MongoDB struct {
		URI string `envconfig:"MONGODB_URI"`
	}
	IntentService struct {
		URL    string `envconfig:"INTENT_SERVICE_URL"`
		APIKey string `envconfig:"INTENT_SERVICE_API_KEY"`
	}
}

// LoadEnv load config variables into Specification
func LoadEnv() (*Specification, error) {
	var config Specification
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
