package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolGraph(t *testing.T) {
	sg := NewSymbolGraph([]string{
		"BJS NYC", "CTU BJS", "BKK NYC",
	})
	assert.Equal(t, len(sg.Adj("BJS")), 2)
	assert.Equal(t, len(sg.Adj("BKK")), 1)
}
