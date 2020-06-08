package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedStack(t *testing.T) {
	stack := NewLinkedStack(&node{Value: "1"})

	stack.Push("2")
	stack.Push("3")

	assert.Equal(t, stack.Pop(), "3")
	assert.Equal(t, stack.Pop(), "2")
	assert.Equal(t, stack.Pop(), "1")
	assert.True(t, stack.IsEmpty())
	assert.Nil(t, stack.Pop())
}
