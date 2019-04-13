package main

import (
	"fmt"
	"github.com/pkg/errors"
	"gotest.tools/assert"
	"net/http"
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
	args       = []string{"run", "."}
	testArgs   = append(args, exampleDir)
)

func TestMain(m *testing.M) {
	s, err := InitTestServer()

	if err != nil {
		panic(errors.Wrap(err, "initialize test server"))
	}

	if err := s.Listen(); err != nil {
		panic(errors.Wrap(err, "start test server"))
	}

	go func () {
		if err := s.Serve(); err != http.ErrServerClosed {
			panic(errors.Wrap(err, "serve test server"))
		}
	}()

	m.Run()

	if err := s.Close(); err != nil {
		panic(errors.Wrap(err, "close test server"))
	}
}

func TestCLI(t *testing.T) {
	c := exec.Command(goExecPath, testArgs...)
	c.Env = os.Environ()

	b, err := c.CombinedOutput()

	fmt.Printf("-- Test output --\n%s\n-- end --\n", string(b))
	assert.NilError(t, err)
}
