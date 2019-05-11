package main

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/config"
	"github.com/robojones/gqltest/example/server"
	"github.com/robojones/gqltest/test_util/tempdir"
	"gotest.tools/assert"
	"os"
	"os/exec"
	"path"
	"runtime"
	"testing"
)

const (
	testPort   = "4444"
	exampleDir = "example"
)

var (
	goExecPath = path.Join(runtime.GOROOT(), "bin/go")
	install    = []string{"install", "."}
	command    = "gqltest"
)

func installCLI(t *testing.T) {
	t.Helper()
	inst := exec.Command(goExecPath, install...)
	inst.Env = os.Environ()
	b, err := inst.CombinedOutput()
	assert.NilError(t, errors.Wrapf(err, "Install CLI\n--- output ---\n%s--- end ---\n", b))
}

func startTestServer(t *testing.T) *server.Server {
	t.Helper()

	if err := os.Setenv("PORT", testPort); err != nil {
		assert.NilError(t, errors.Wrap(err, "set PORT env for test server"))
	}
	s := InitTestServer()
	if err := os.Unsetenv("PORT"); err != nil {
		assert.NilError(t, errors.Wrap(err, "unset PORT env for test server"))
	}

	go func() {
		assert.NilError(t, s.Listen())
	}()

	return s
}

func TestCLI(t *testing.T) {
	s := startTestServer(t)
	defer func() {
		assert.NilError(t, s.Close())
	}()

	installCLI(t)

	c := exec.Command(command)
	c.Dir = exampleDir
	c.Env = os.Environ()
	b, err := c.CombinedOutput()

	assert.NilError(t, errors.Wrapf(err, "-- Test output --\n%s\n-- end --\n", string(b)))
}

func TestInitTesterErrorConfigMissing(t *testing.T) {
	d := tempdir.Create(t)
	defer tempdir.Remove(t, d)

	_, e := InitTester(config.WorkinDirectoryName(d))

	assert.Assert(t, os.IsNotExist(errors.Cause(e)))
}

func TestInitTesterErrorEndpointMissing(t *testing.T) {
	d := tempdir.Create(t)
	defer tempdir.Remove(t, d)

	tempdir.File(t, d, "gqltest.yml", "")

	_, e := InitTester(config.WorkinDirectoryName(d))
	assert.Equal(t, errors.Cause(e), config.VerificationError)
}
