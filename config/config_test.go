package config

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewDefaultConfig(t *testing.T) {
	c := newDefaultConfig()

	assert.Equal(t, c.Endpoint, "")
	assert.Equal(t, c.TestRoot, DefaultTestRoot)
}
