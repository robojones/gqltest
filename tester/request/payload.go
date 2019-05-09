package request

import (
	"bytes"
	"encoding/json"
)

type Variables map[string]interface{}

func NewVariables() Variables {
	return make(Variables)
}

type Payload struct {
	OperationName string
	Query         string
	Variables     Variables
}

func NewPayload(operation string, query string, vars Variables) *Payload {
	return &Payload{
		OperationName: operation,
		Query:         query,
		Variables:     vars,
	}
}

func (p *Payload) Body() (*bytes.Reader, error) {
	b, err := json.Marshal(p)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}
