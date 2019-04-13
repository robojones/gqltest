package config

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

const testEndpoint = "http://testendpoint.test"
const testTestRoot = "testtestroot"

var testConfigContent = fmt.Sprintf("endpoint: %s\ntestRoot: %s", testEndpoint, testTestRoot)

func TestParseConfig(t *testing.T) {
	const testEndpoint = "http://testendpoint.test"
	const testTestRoot = "testtestroot"

	b := testConfigContent
	c, err := parseConfig([]byte(b))

	assert.NilError(t, err)
	assert.Equal(t, c.Endpoint, testEndpoint)
	assert.Equal(t, c.TestRoot, testTestRoot)
}
