package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(10)
	uf.union(4, 3)
	uf.union(3, 8)
	uf.union(6, 5)
	uf.union(9, 4)
	uf.union(2, 1)
	uf.union(5, 0)
	uf.union(7, 2)
	uf.union(6, 1)

	assert.True(t, uf.connected(2, 6))
	assert.True(t, uf.connected(3, 9))
	assert.Equal(t, uf.unionCount(), 2)
}

func TestForestUnionFind(t *testing.T) {
	uf := NewTreeUnionFind(10)
	uf.union(4, 3)
	uf.union(3, 8)
	uf.union(6, 5)
	uf.union(9, 4)
	uf.union(2, 1)
	uf.union(5, 0)
	uf.union(7, 2)
	uf.union(6, 1)

	assert.True(t, uf.connected(2, 6))
	assert.True(t, uf.connected(3, 9))
	assert.True(t, uf.connected(5, 1))
	assert.Equal(t, uf.unionCount(), 2)
}
