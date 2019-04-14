package config

import (
	"github.com/robojones/gqltest/test_util/tempdir"
	"gotest.tools/assert"
	"path"
	"testing"
)

func TestNewConfig(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	tempdir.File(t, dir, ConfigFileName, testConfigContent)

	c, err := NewConfig(WD(dir))

	assert.NilError(t, err)
	assert.Equal(t, c.Endpoint(), testEndpoint)
	assert.Equal(t, c.TestRoot(), path.Join(dir, testTestRoot))
}