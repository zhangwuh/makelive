package structure

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjGraph(t *testing.T) {
	g := NewAdjGraph(13, []edge{
		{0, 5}, {4, 3}, {0, 1}, {9, 12}, {6, 4}, {5, 4}, {0, 2}, {11, 12}, {9, 10}, {0, 6},
		{7, 8}, {9, 11}, {5, 3},
	})

	assert.Equal(t, g.V(), 13)
	assert.Equal(t, g.E(), 13)
	resolver := NewGraphResolver(g)
	assert.Equal(t, resolver.Degree(0), 4)
	assert.Equal(t, resolver.MaxDegree(), 4)
}

func TestUnionFindSearcher(t *testing.T) {
	g := NewAdjGraph(13, []edge{
		{0, 5}, {4, 3}, {0, 1}, {9, 12}, {6, 4}, {5, 4}, {0, 2}, {11, 12}, {9, 10}, {0, 6},
		{7, 8}, {9, 11}, {5, 3},
	})

	searcher := newUfSearcher(g)
	assert.Equal(t, searcher.Count(5), 6)
	assert.Equal(t, searcher.Count(9), 3)
	assert.Equal(t, searcher.Count(7), 1)
	assert.True(t, searcher.Marked(0, 5))
	assert.True(t, searcher.Marked(5, 0))
	assert.True(t, searcher.Marked(9, 12))
	assert.True(t, searcher.Marked(3, 0))

	assert.False(t, searcher.Marked(12, 1))
}

func TestDfsSearcher(t *testing.T) {
	g := NewAdjGraph(13, []edge{
		{0, 5}, {4, 3}, {0, 1}, {6, 4}, {5, 4}, {0, 2}, {11, 12}, {9, 10}, {0, 6},
		{7, 8}, {9, 11}, {5, 3},
	})

	searcher := newDfsSearcher(g)
	assert.Equal(t, searcher.Count(5), 6)
	assert.Equal(t, searcher.Count(9), 3)
	assert.Equal(t, searcher.Count(7), 1)
	assert.True(t, searcher.Marked(0, 5))
	assert.True(t, searcher.Marked(5, 0))
	assert.True(t, searcher.Marked(9, 12))
	assert.True(t, searcher.Marked(3, 0))

	assert.False(t, searcher.Marked(12, 1))

	fmt.Println(searcher.pathTo(0, 3))
	fmt.Println(searcher.pathTo(7, 8))
	fmt.Println(searcher.pathTo(9, 12))
	fmt.Println(searcher.pathTo(0, 12))
}

func TestBfsSearcher(t *testing.T) {
	g := NewAdjGraph(13, []edge{
		{0, 5}, {4, 3}, {0, 1}, {9, 12}, {6, 4}, {5, 4}, {0, 2}, {11, 12}, {9, 10}, {0, 6},
		{7, 8}, {9, 11}, {5, 3},
	})

	searcher := newBfsSearcher(g)
	assert.Equal(t, searcher.Count(5), 6)
	assert.Equal(t, searcher.Count(9), 3)
	assert.Equal(t, searcher.Count(7), 1)
	assert.True(t, searcher.Marked(0, 5))
	assert.True(t, searcher.Marked(5, 0))
	assert.True(t, searcher.Marked(9, 12))
	assert.True(t, searcher.Marked(3, 0))

	assert.False(t, searcher.Marked(12, 1))

	fmt.Println(searcher.pathTo(0, 3))
	fmt.Println(searcher.pathTo(7, 8))
	fmt.Println(searcher.pathTo(9, 12))
	fmt.Println(searcher.pathTo(0, 12))
}
