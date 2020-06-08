package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	for _, v := range res {
		fmt.Println(v)
	}

	nums = []int{0, 0, 0}
	res = threeSum(nums)
	for _, v := range res {
		fmt.Println(v)
	}
}
