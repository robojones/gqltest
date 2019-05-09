package testerror

import (
	"gotest.tools/assert"
	"testing"
)

func TestHighlight(t *testing.T) {
	s := highlight(testDirective)

	assert.Equal(t, s, "2| query Test {\n"+
		"3>   wazzup "+exampleDirective+"\n"+
		"4| }")
}

func TestFormatLineNumber(t *testing.T) {
	assert.Equal(t, formatLineNumber(1, 1, false), "1| ")
	assert.Equal(t, formatLineNumber(10, 4, false), "  10| ")
	assert.Equal(t, formatLineNumber(1, 1, true), "1> ")
}

func TestGetNumberLength(t *testing.T) {
	assert.Equal(t, getNumberLength(10), 2)
	assert.Equal(t, getNumberLength(0), 1)
	assert.Equal(t, getNumberLength(1000), 4)
}
