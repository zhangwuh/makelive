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
	g := NewDiAdjGraph(7, []edge{
		{0, 1}, {2, 3}, {0, 3}, {4, 0}, {1, 4}, {5, 6},
	})
	dfs := NewDirectedDfs(g, []int{0, 5})
	assert.True(t, dfs.isMarked(1))
	assert.True(t, dfs.isMarked(3))
	assert.True(t, dfs.isMarked(4))
	assert.False(t, dfs.isMarked(2))
	assert.True(t, dfs.isMarked(6))
}

func TestNewDirectedCycleDetector(t *testing.T) {
	g := NewDiAdjGraph(7, []edge{
		{0, 1}, {2, 3}, {0, 3}, {4, 0}, {1, 4}, {5, 6},
	})
	dfs := NewDirectedCycleDetector(g)
	assert.True(t, dfs.HasCycle())

	g = NewDiAdjGraph(4, []edge{
		{0, 1}, {0, 2}, {1, 3}, {2, 3},
	})
	dfs = NewDirectedCycleDetector(g)
	assert.False(t, dfs.HasCycle())
}

func Test_DepthFirstOrder(t *testing.T) {
	g := NewDiAdjGraph(13, []edge{
		{0, 5}, {0, 1}, {0, 6}, {2, 0}, {2, 3}, {3, 5}, {5, 4}, {6, 4}, {7, 6}, {8, 7}, {6, 9}, {9, 10}, {9, 11}, {11, 12}, {9, 12},
	})

	df := NewDepthFirstOrder(g)

	fmt.Println(df.pre)
	fmt.Println(df.post)
	fmt.Println(df.reversePostQueue())
}

func TestNewTopologicalSorter(t *testing.T) {
	g := NewDiAdjGraph(13, []edge{
		{0, 5}, {0, 1}, {0, 6}, {2, 0}, {2, 3}, {3, 5}, {5, 4}, {6, 4}, {7, 6}, {8, 7}, {6, 9}, {9, 10}, {9, 11}, {11, 12}, {9, 12},
	})
	sorter := NewTopologicalSorter(g)
	fmt.Println(sorter.order())
}

func TestNormalSCC(t *testing.T) {
	g := NewDiAdjGraph(13, []edge{
		{0, 1}, {1, 2}, {2, 0}, {1, 6}, {6, 0}, {2, 3}, {3, 4}, {4, 5}, {5, 3},
	})
	scc := NewNormalSCC(g)
	assert.True(t, scc.Connected(2, 6))
	assert.True(t, scc.Connected(3, 5))
	assert.False(t, scc.Connected(2, 3))
}
