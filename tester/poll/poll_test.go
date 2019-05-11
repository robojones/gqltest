package poll

import (
	"github.com/pkg/errors"
	"github.com/robojones/gqltest/test_util/server"
	"gotest.tools/assert"
	"net"
	"testing"
	"time"
)

func TestPoll(t *testing.T) {
	s := server.ServeJSON(t, "{}")
	defer s.Close()

	err := Poll(s.URL, 5*time.Second)
	assert.NilError(t, err)
}

func TestPollTimeout(t *testing.T) {
	s := server.ServeJSON(t, "{}")
	defer s.Close()
	err := Poll(s.URL, 1)
	assert.Assert(t, errors.Cause(err).(net.Error).Timeout())
}
