package main

import (
	"fmt"
	"github.com/pkg/errors"
	"gotest.tools/assert"
	"log"
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
	install    = []string{"install", "."}
	command    = "gqltest"
)

func TestMain(m *testing.M) {
	if err := os.Setenv("PORT", testPort); err != nil {
		panic(errors.Wrap(err, "set PORT env for test server"))
	}

	s, err := InitTestServer()

	if err != nil {
		panic(errors.Wrap(err, "initialize test server"))
	}

	if err := s.Listen(); err != nil {
		panic(errors.Wrap(err, "start test server"))
	}

	log.Printf("test server online")

	go func() {
		log.Printf("serve")
		if err := s.Serve(); err != http.ErrServerClosed {
			log.Printf("server closed")
			panic(errors.Wrap(err, "serve test server"))
		}
	}()

	inst := exec.Command(goExecPath, install...)
	inst.Env = os.Environ()
	b, err := inst.CombinedOutput()

	if err != nil {
		panic(errors.Errorf(
			"error during installation \"%s\"\n-- output --\n%s-- end --",
			err.Error(),
			string(b),
		))
	}

	m.Run()

	log.Printf("close test server")

	if err := s.Close(); err != nil {
		panic(errors.Wrap(err, "close test server"))
	}
}

func TestCLI(t *testing.T) {
	c := exec.Command(command)
	c.Dir = exampleDir
	c.Env = os.Environ()

	b, err := c.CombinedOutput()

	fmt.Printf("-- Test output --\n%s\n-- end --\n", string(b))
	assert.NilError(t, err)
}
