package testerror

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func stringifyValue(v interface{}) string {
	r, err := json.Marshal(v)

	if err != nil {
		panic(errors.WithStack(err))
	}

	return string(r) + string(NEWLINE)
}

func stringifyObject(v interface{}) string {
	r, err := json.MarshalIndent(v, Prefix, Indent)

	if err != nil {
		panic(errors.WithStack(err))
	}

	return string(r) + string(NEWLINE)
}
