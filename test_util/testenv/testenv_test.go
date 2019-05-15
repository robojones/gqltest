package testenv

import (
	"gotest.tools/assert"
	"os"
	"testing"
)

const (
	origVal = "origVal"
	testKey = "testenv_key"
	testVal = "testenv_val"
)

func TestSet(t *testing.T) {
	Set(t, testKey, testVal)
	assert.Equal(t, os.Getenv(testKey), testVal)
}

func TestUnset(t *testing.T) {
	Set(t, testKey, testVal)
	assert.Equal(t, os.Getenv(testKey), testVal)

	Unset(t, testKey)
	_, ok := os.LookupEnv(testKey)
	assert.Assert(t, !ok)
}

func TestEnvVersion_Reset(t *testing.T) {
	assert.NilError(t, os.Setenv(testKey, origVal))

	orig := Set(t, testKey, testVal)
	assert.Equal(t, orig.ok, true)
	assert.Equal(t, orig.key, testKey)
	assert.Equal(t, orig.val, origVal)

	orig.Reset()
	assert.Equal(t, os.Getenv(testKey), origVal)
}

func TestEnvVersion_ResetUnset(t *testing.T) {
	assert.NilError(t, os.Unsetenv(testKey))

	v := Set(t, testKey, testVal)
	assert.Equal(t, v.ok, false)
	assert.Equal(t, v.key, testKey)
	assert.Equal(t, os.Getenv(testKey), testVal)

	v.Reset()
	_, ok := os.LookupEnv(testKey)
	assert.Assert(t, !ok)
}
