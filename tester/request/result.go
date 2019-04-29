package request

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type ErrorLocation struct {
	Row    int
	Column int
}

type ErrorExtensions = map[string]interface{}

type Error struct {
	Message    string
	Locations  []*ErrorLocation
	Extensions ErrorExtensions
}

type Result interface{}

func ParseResult(body []byte) (*Result, error) {
	result := new(Result)
	err := json.Unmarshal(body, result)

	if err != nil {
		return nil, errors.Wrapf(err, "parse body \"%s\"", string(body))
	}

	return result, nil
}
