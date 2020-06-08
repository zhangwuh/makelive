package structure

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiAdjGraph(t *testing.T) {
	g := NewDiAdjGraph(5, []edge{
		{0, 1}, {2, 3}, {0, 3}, {4, 0}, {1, 4},
	})

	assert.Equal(t, g.V(), 5)
	assert.Equal(t, g.E(), 5)
	resolver := NewGraphResolver(g.Reverse())
	fmt.Println(resolver.Desc())
}

func TestNewDiAdjGraph(t *testing.T) {
	g := NewDiAdjGraph(5, []edge{
		{0, 1}, {2, 3}, {0, 3}, {4, 0}, {1, 4},
	})
	dfs := NewDirectedDfs(g, []int{0})
	assert.True(t, dfs.isMarked(1))
	assert.True(t, dfs.isMarked(3))
	assert.True(t, dfs.isMarked(4))
	assert.False(t, dfs.isMarked(2))
}
