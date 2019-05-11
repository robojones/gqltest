package config

import (
	"fmt"
	"github.com/robojones/gqltest/test_util/tempdir"
	"gotest.tools/assert"
	"testing"
	"time"
)

const (
	testEndpoint           = "http://testendpoint.test"
	testSchemaGlob         = "someschema/*.graphqls"
	testTestGlob           = "some/*.graphql"
	testStartTimeout int64 = 1000 // 1s
	testTestTimeout  int64 = 500  // 0.5s
)

var testConfigContent = fmt.Sprintf(`
endpoint: %s
schema: %s
tests: %s
startTimeout: %d
testTimeout: %d
`, testEndpoint, testSchemaGlob, testTestGlob, testStartTimeout, testTestTimeout)

func TestNewConfig(t *testing.T) {
	dir := tempdir.Create(t)
	defer tempdir.Remove(t, dir)

	tempdir.File(t, dir, ConfigFileName, testConfigContent)

	c, err := NewConfig(WorkinDirectoryName(dir))

	assert.NilError(t, err)
	assert.Equal(t, c.Endpoint(), testEndpoint)
	assert.Equal(t, c.SchemaGlob(), testSchemaGlob)
	assert.Equal(t, c.TestGlob(), testTestGlob)
	assert.Equal(t, c.StartTimeout(), time.Duration(testStartTimeout) * time.Millisecond)
	assert.Equal(t, c.TestTimeout(), time.Duration(testTestTimeout) * time.Millisecond)
}
