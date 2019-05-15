package directives

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/test_util/testenv"
	"gotest.tools/assert"
	"testing"
)

func TestGetBranch(t *testing.T) {
	const testBranch = "test_branch"

	env := testenv.Set(t, BranchEnvKey, testBranch)
	defer env.Reset()

	b := getBranch()
	assert.Equal(t, b, testBranch)
}

func TestGetBranchDefault(t *testing.T) {
	env := testenv.Unset(t, BranchEnvKey)
	defer env.Reset()

	b := getBranch()
	assert.Equal(t, b, DefaultBranch)
}

func TestGet(t *testing.T) {
	d, err := Get()
	assert.NilError(t, err)
	assert.Assert(t, d.Input != "")
}

func TestGetStatusError(t *testing.T) {
	const unknownBranch = "asdfhlkwlerkj"
	env := testenv.Set(t, BranchEnvKey, unknownBranch)
	defer env.Reset()

	_, err := Get()

	stErr := errors.Cause(err).(*StatusError)
	assert.Equal(t, stErr.status, "404 Not Found")
}
