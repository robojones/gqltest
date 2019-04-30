package json

import (
	"encoding/json"
	"github.com/pkg/errors"
)

const (
	Prefix = ""
	Indent = "  "
)

func StringifyValue(v interface{}) string {
	r, err := json.Marshal(v)

	if err != nil {
		panic(errors.WithStack(err))
	}

	return string(r)
}

func StringifyObject(v interface{}) string {
	r, err := json.MarshalIndent(v, Prefix, Indent)

	if err != nil {
		panic(errors.WithStack(err))
	}

	return string(r)
}
