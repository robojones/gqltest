package config

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewDefaultConfigData(t *testing.T) {
	c := newDefaultConfigData()

	assert.Equal(t, c.Endpoint, "")
	assert.Equal(t, c.SchemaGlob, DefaultSchemaGlob)
	assert.Equal(t, c.TestGlob, DefaultTestGlob)
	assert.Equal(t, c.StartTimeout, DefaultStartTimeout)
	assert.Equal(t, c.TestTimeout, DefaultTestTimeout)
}

