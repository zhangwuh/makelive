package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	assert.True(t, isPalindrome(121))
	assert.False(t, isPalindrome(-121))
	assert.False(t, isPalindrome(10))
}
