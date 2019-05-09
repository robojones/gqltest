package config

import (
	"github.com/robojones/gqltest/test_util/tempdir"
	"gotest.tools/assert"
	"path"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	tempdir.File(t, dir, ConfigFileName, testConfigContent)

	c, err := NewConfig(WorkinDirectoryName(dir))

	assert.NilError(t, err)
	assert.Equal(t, c.Endpoint(), testEndpoint)
	assert.Equal(t, c.TestRoot(), path.Join(dir, testTestRoot))
	assert.Equal(t, c.StartTimeout(), time.Duration(testStartTimeout) * time.Millisecond)
	assert.Equal(t, c.TestTimeout(), time.Duration(testTestTimeout) * time.Millisecond)
}
