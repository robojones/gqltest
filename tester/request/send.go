package request

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const contentType = "text/json"

func Send(endpoint string, r *Payload) (Result, error) {
	b, err := r.Body()

	if err != nil {
		return nil, errors.Wrap(err, "stringify body")
	}

	response, err := http.Post(endpoint, contentType, b)

	if err != nil {
		return nil, errors.Wrap(err, "post request")
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, errors.Errorf("Expected 2xx Error but got %s", response.Status)
	}

	resBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.New("Error reading body.")
	}

	return ParseResult(resBody)
}
