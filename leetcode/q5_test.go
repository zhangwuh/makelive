package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, longestPalindrome("babad"), "bab")
	assert.Equal(t, longestPalindrome("cbbd"), "bb")
	assert.Equal(t, longestPalindrome("a"), "a")
	assert.Equal(t, longestPalindrome(""), "")
	assert.Equal(t, longestPalindrome("abb"), "bb")
}
