package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))
	assert.Equal(t, 0, reverse(1323423233))
	assert.Equal(t, 0, reverse(-1323423233))
	assert.Equal(t, 2147483641, reverse(1463847412))
	assert.Equal(t, -2147483641, reverse(-1463847412))
}
