package request

import (
	"encoding/json"
	"github.com/robojones/gqltest/test_util/server"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	expectedResponse := "{\"this is json\":true}"
	s := server.ServeJSON(t, expectedResponse)
	defer s.Close()

	p := NewPayload("test something", "query { some }", nil)
	r, err := Send(s.URL, p, time.Second*10)
	assert.NilError(t, err)

	bod, err := json.Marshal(r)
	assert.NilError(t, err)
	assert.Equal(t, string(bod), expectedResponse)

}
