package request

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewVariables(t *testing.T) {
	vars := NewVariables()

	assert.Assert(t, vars != nil)
}

func TestNewPayload(t *testing.T) {
	const (
		op    = "test operation"
		query = "query TestOperation { username }"
	)

	vars := NewVariables()

	p := NewPayload(op, query, vars)

	assert.Equal(t, p.OperationName, op)
	assert.Equal(t, p.Query, query)
	assert.DeepEqual(t, p.Variables, vars)
}

func TestPayload_Body(t *testing.T) {

}
