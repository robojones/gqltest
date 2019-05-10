package server

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type staticJSONHandler struct {
	res string
}

func (s *staticJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Add("Content-Type", "test/json")

	_, err := fmt.Fprint(w, s.res)
	if err != nil {
		panic(errors.WithStack(err))
	}
}

func ServeJSON(t *testing.T, response string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(&staticJSONHandler{res: response})
}
