package config

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

const (
	testEndpoint           = "http://testendpoint.test"
	testTestRoot           = "testtestroot"
	testStartTimeout int64 = 1000 // 1s
	testTestTimeout  int64 = 500  // 0.5s
)

var testConfigContent = fmt.Sprintf(`
endpoint: %s
testRoot: %s
startTimeout: %d
testTimeout: %d
`, testEndpoint, testTestRoot, testStartTimeout, testTestTimeout)

func TestParseConfig(t *testing.T) {
	const testEndpoint = "http://testendpoint.test"
	const testTestRoot = "testtestroot"

	b := testConfigContent
	c, err := parseConfig([]byte(b))

	assert.NilError(t, err)
	assert.Equal(t, c.Endpoint, testEndpoint)
	assert.Equal(t, c.TestRoot, testTestRoot)
	assert.Equal(t, c.StartTimeout, testStartTimeout)
	assert.Equal(t, c.TestTimeout, testTestTimeout)
}
