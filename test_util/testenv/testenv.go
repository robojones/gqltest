package testenv

import (
	"gotest.tools/assert"
	"os"
	"testing"
)

type EnvVersion struct {
	t   *testing.T
	key string
	ok  bool
	val string
}

func Set(t *testing.T, key string, val string) *EnvVersion {
	t.Helper()

	v, ok := os.LookupEnv(key)
	orig := &EnvVersion{
		t:   t,
		key: key,
		ok:  ok,
		val: v,
	}

	assert.NilError(t, os.Setenv(key, val))

	return orig
}

func Unset(t *testing.T, key string) *EnvVersion {
	t.Helper()

	v, ok := os.LookupEnv(key)
	orig := &EnvVersion{
		t:   t,
		key: key,
		ok:  ok,
		val: v,
	}

	assert.NilError(t, os.Unsetenv(key))

	return orig
}

func (e *EnvVersion) Reset() {
	e.t.Helper()

	if e.ok {
		assert.NilError(e.t, os.Setenv(e.key, e.val))
	} else {
		assert.NilError(e.t, os.Unsetenv(e.key))
	}
}
