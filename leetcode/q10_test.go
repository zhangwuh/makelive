package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isMatch(t *testing.T) {
	assert.False(t, isMatch2("aa", "a"))
	assert.True(t, isMatch2("aa", "a*"))
	assert.True(t, isMatch2("ab", ".*"))
	assert.True(t, isMatch2("aab", "c*a*b"))
	assert.True(t, isMatch2("", "c*a*"))
	assert.False(t, isMatch2("mississippi", "mis*is*p*."))
	assert.True(t, isMatch2("aaa", "ab*a*c*a"))
	assert.False(t, isMatch2("aaa", "aaaa"))
	assert.False(t, isMatch2("a", ".*..a*"))
	assert.True(t, isMatch2("ab", ".*.."))
}

func TestDeque(t *testing.T) {
	dq := &deque{}
	dq.Push("a")
	dq.Push("b")
	dq.Push("c")
	dq.Add("e")
	fmt.Println(dq.Pop())
	fmt.Println(dq.Offer())
	fmt.Println(dq.Pop())
	fmt.Println(dq.Offer())
}

func Test_resolveMatchers(t *testing.T) {
	m := resolveMatchers("ab*a*c*a")
	fmt.Println(m)
	f := m.Pop().(matcher)
	x, y, _ := f.f('a')
	fmt.Println(x)
	fmt.Println(y)
}
