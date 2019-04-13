package config

import (
	"github.com/robojones/gqltest/test_util/tempdir"
	"github.com/robojones/gqltest/test_util/tempfile"
	"gotest.tools/assert"
	"testing"
)

func TestReadConfigData(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	tempfile.Create(t, dir, ConfigFileName, testConfigContent)

	c, err := readConfigData(WD(dir))

	assert.NilError(t, err)
	assert.Equal(t, c.Endpoint, testEndpoint)
	assert.Equal(t, c.TestRoot, testTestRoot)
}
