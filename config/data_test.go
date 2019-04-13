package config

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewDefaultConfigData(t *testing.T) {
	c := newDefaultConfigData()

	assert.Equal(t, c.Endpoint, "")
	assert.Equal(t, c.TestRoot, DefaultTestRoot)
}

