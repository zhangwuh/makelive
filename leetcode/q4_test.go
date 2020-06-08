package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Merge(t *testing.T) {
	n1 := []int{1, 2}
	n2 := []int{3, 4}
	assert.Equal(t, findMedianSortedArrays(n1, n2), 2.5)

	n1 = []int{1, 3, 5, 7, 9, 11}
	n2 = []int{0, 2, 4, 6, 8}

	assert.Equal(t, findMedianSortedArrays(n1, n2), float64(5))

	n1 = []int{0, 0}
	n2 = []int{0, 0}

	assert.Equal(t, findMedianSortedArrays(n1, n2), float64(0))
}
