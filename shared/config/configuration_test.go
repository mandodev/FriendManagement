package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfigurationFromFile(t *testing.T) {
	cfg, err := New()

	assert.Empty(t, err)

	configuration := *cfg

	// test to configuration object
	assert.NotEmpty(t, configuration)

	//test config  value
	assert.Equal(t, configuration.Server.Mode, "debug")
	assert.Equal(t, configuration.Server.Addr, ":8080")
	assert.Equal(t, configuration.Server.ShutdownTimeout, 5)
}
