package server

import (
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestServeJSON(t *testing.T) {
	expectedResponse := "{\"this is json\":true}"
	s := ServeJSON(t, expectedResponse)
	defer s.Close()

	r, err := http.Post(s.URL, "text/json", nil)
	assert.NilError(t, err)

	resp, err := ioutil.ReadAll(r.Body)
	assert.NilError(t, err)
	assert.Equal(t, string(resp), expectedResponse)
}
