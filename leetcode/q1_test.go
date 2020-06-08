package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQ1(t *testing.T) {
	input := []int{7, 2, 3, 5, 4, 15, 11}
	r := twoSumOnMap(input, 11)
	assert.Equal(t, input[r[0]]+input[r[1]], 11)

}
