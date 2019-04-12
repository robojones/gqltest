package request

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const (
	contentType = "text/json"
	url         = "http://localhost:4000/query"
)

type Request struct {
	payload *Payload
}

func NewRequest(payload *Payload) *Request {
	return &Request{
		payload: payload,
	}
}

func (r *Request) Body() (*bytes.Reader, error) {
	b, err := json.Marshal(r.payload)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

func Send(r *Request) (*Result, error) {
	b, err := r.Body()

	if err != nil {
		return nil, errors.Wrap(err, "stringify body")
	}

	response, err := http.Post(url, contentType, b)

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

	result := new(Result)
	err = json.Unmarshal(resBody, result)

	if err != nil {
		return nil, errors.Wrapf(err, "parse body \"%s\"", string(resBody))
	}

	return result, nil
}
