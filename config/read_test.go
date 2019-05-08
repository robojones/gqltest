package config

import (
	"github.com/robojones/gqltest/test_util/tempdir"
	"gotest.tools/assert"
	"testing"
)

func TestReadConfigData(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	tempdir.File(t, dir, ConfigFileName, testConfigContent)

	c, err := readConfigData(WorkinDirectoryName(dir))

	assert.NilError(t, err)
	assert.Equal(t, c.Endpoint, testEndpoint)
	assert.Equal(t, c.TestRoot, testTestRoot)
}
