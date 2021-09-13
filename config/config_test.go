package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertConfig(t *testing.T, globalconfig *Specification) {
	assert.Equal(t, 8080, globalconfig.HTTP.Port)
	assert.Equal(t, "http://intent-service.ai", globalconfig.IntentService.URL)
	assert.Equal(t, "localhost", globalconfig.MongoDB.URI)
}

func TestLoadEnv(t *testing.T) {
	// Global Config
	setGlobalConfigEnv()

	globalConfig, err := LoadEnv()
	assert.Nil(t, err)

	assertConfig(t, globalConfig)
}

func setGlobalConfigEnv() {
	os.Setenv("HTTP_PORT", "8080")
	os.Setenv("INTENT_SERVICE_URL", "http://intent-service.ai")
	os.Setenv("MONGODB_URI", "localhost")
}
