package test

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewTest(t *testing.T) {
	doc := getTestDoc()
	test := NewTest(doc)

	assert.Assert(t, test != nil)
}
