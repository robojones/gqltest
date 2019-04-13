package config

import (
	"github.com/pkg/errors"
	"gotest.tools/assert"
	"testing"
)

func TestVerify(t *testing.T) {
	c := newDefaultConfig()
	c.Endpoint = "http://localhost:2345"

	err := verifyConfig(c)

	assert.NilError(t, err)
}

func TestVerify_Fail(t *testing.T) {
	c := newDefaultConfig()

	err := verifyConfig(c)

	assert.Equal(t, errors.Cause(err), VerificationError)
}
