package request

import (
	"encoding/json"
	"fmt"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type staticJSONHandler struct {
	t *testing.T
	res string
}

func (s *staticJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.t.Helper()

	h := w.Header()
	h.Add("Content-Type", "test/json")

	_, err := fmt.Fprint(w, s.res)

	assert.NilError(s.t, err)
}

func ServeJSON(t *testing.T, response string) *httptest.Server {
	return httptest.NewServer(&staticJSONHandler{t: t, res: response})
}

func TestSend(t *testing.T) {
	expectedResponse := "{\"this is json\":true}"
	server := ServeJSON(t, expectedResponse)
	defer server.Close()

	p := NewPayload("test something", "query { some }", nil)
	r, err := Send(server.URL, p, time.Second * 10)
	assert.NilError(t, err)

	bod, err := json.Marshal(r)
	assert.NilError(t, err)
	assert.Equal(t, string(bod), expectedResponse)

}
