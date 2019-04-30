package json

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func Value(s string) interface{} {
	v := new(interface{})
	if err := json.Unmarshal([]byte(s), v); err != nil {
		panic(errors.WithStack(err))
	}
	return *v
}
