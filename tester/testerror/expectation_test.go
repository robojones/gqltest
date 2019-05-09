package testerror

import (
	"encoding/json"
	"github.com/pkg/errors"
	json2 "github.com/robojones/gqltest/tester/json"
	"gotest.tools/assert"
	"log"
	"strings"
	"testing"
)

const wrappedMessage = "wazzup"

var (
	orig = errors.New(wrappedMessage)
	path = []string{"cool", "path"}
	res  = `{ "cool": { "path": "wazzup" }}`
	r    = new(interface{})
)

func init() {
	err := json.Unmarshal([]byte(res), r)
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewExpectationError(t *testing.T) {
	e := NewExpectationError(orig, path, r)

	assert.ErrorType(t, e, new(ExpectationError))
	assert.Equal(t, e.err, orig)
	assert.Equal(t, e.result, r)
	assert.DeepEqual(t, e.path, path)
}

func TestExpectationError_Cause(t *testing.T) {
	e := NewExpectationError(orig, path, r)

	assert.Equal(t, e.Cause(), orig)
}

func TestExpectationError_Error(t *testing.T) {
	e := NewExpectationError(orig, path, r)

	assert.ErrorContains(t, e, json2.StringifyObject(r))
	assert.ErrorContains(t, e, strings.Join(path, "."))
	assert.ErrorContains(t, e, orig.Error())
}
