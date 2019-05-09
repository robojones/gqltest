package test

import (
	"gotest.tools/assert"
	"testing"
)

func TestTest_Operation(t *testing.T) {
	doc := getTestDoc()
	test := NewTest(doc)

	assert.Equal(t, test.Operation(), test.Document.Operations[0])
}
