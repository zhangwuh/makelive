package practise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloomFilter(t *testing.T) {
	filter := NewBloomFilter(0.001, 1000)
	filter.Set([]byte("abcd"))
	filter.Set([]byte("aaaa"))
	filter.Set([]byte("bbdcd"))

	assert.Equal(t, filter.Get([]byte("abcd")), true)
	assert.Equal(t, filter.Get([]byte("bbdcd")), true)

	assert.Equal(t, filter.Get([]byte("abc")), false)
}
