package bodypart

import (
	"gotest.tools/assert"
	"testing"
)

func TestParseSource(t *testing.T) {
	parts, err := ParseSource(testSource)

	assert.NilError(t, err)

	assert.Equal(t, parts[0].Source, testSource)
	assert.DeepEqual(t, *parts[0].node.GetLoc(), *queryNode.GetLoc())
	assert.Equal(t, parts[1].Source, testSource)
	assert.DeepEqual(t, *parts[1].node.GetLoc(), *mutationNode.GetLoc())
}
