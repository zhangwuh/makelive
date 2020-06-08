package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
	assert.Equal(t, -42, myAtoi("    -42"))
	assert.Equal(t, 4193, myAtoi("4193 with words"))
	assert.Equal(t, 0, myAtoi("words and 987"))
	assert.Equal(t, math.MinInt32, myAtoi("-91283472332"))
	assert.Equal(t, 0, myAtoi("+"))
	assert.Equal(t, 2147483646, myAtoi("2147483646"))
}
