package leetcode

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		input string
		res   int
	}{
		{
			"abcabcbb", 3,
		},
		{
			"bbbbb", 1,
		},
		{
			"pwwkew", 3,
		},
		{
			"dvdf", 3,
		},
		{
			"ohomm", 3,
		},
	}
	for _, c := range cases {
		m := lengthOfLongestSubstring(c.input)
		fmt.Println(m)
		assert.Equal(t, m, c.res)
	}

}

func Benchmark_lengthOfLongestSubstring(t *testing.B) {
	sample, _ := ioutil.ReadFile("q3")
	for i := 0; i < t.N; i++ {
		fmt.Println(lengthOfLongestSubstring(string(sample)))
	}
}
