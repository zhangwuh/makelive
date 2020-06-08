package leetcode

import "sort"

const target = 0

func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)
	lastC := target + 1
	for i, c := range nums {
		if c > target {
			return
		}
		if lastC == c {
			continue
		}
		others := scan(nums[i+1:], c)
		if len(others) > 0 {
			res = append(res, others...)
		}
		lastC = c
	}
	return
}

func scan(nums []int, c int) (res [][]int) {
	left := 0
	right := len(nums) - 1
	for {
		if left >= right {
			return
		}
		leftVal := nums[left]
		rightVal := nums[right]
		cr := compareToTarget(leftVal, rightVal, c)
		if cr == 0 {
			res = append(res, []int{c, leftVal, rightVal})
			left = skip(nums, left, 1)
			right = skip(nums, right, -1)
		} else if cr == -1 {
			left++
		} else {
			right--
		}
	}

}

func skip(nums []int, pos int, step int) int {
	current := nums[pos]
	for pos < len(nums) {
		pos += step
		if pos < len(nums) && pos >= 0 && nums[pos] == current {
			continue
		}
		break
	}
	return pos
}

func compareToTarget(a, b, c int) int {
	res := a + b + c
	if res == target {
		return 0
	}
	if res < target {
		return -1
	}
	return 1
}
