// q1.go  * Created on  2020/5/26
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
package leetcode

import (
	"math/rand"
	"time"
)

type node struct {
	index int
	value int
}

func twoSumOnMap(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	res := make([]int, 0)
	for i, v := range nums {
		if t, ok := m[target-v]; ok {
			res = append(res, i)
			res = append(res, t)
			break
		}
		m[v] = i
	}
	return res
}

func twoSum(nums []int, target int) []int {
	var nodes []*node
	for i, v := range nums {
		nodes = append(nodes, &node{i, v})
	}
	qsort(nodes)
	for i, tom := range nodes {
		swap(nodes, 0, i)
		j := binarySearch(nodes[1:], target-tom.value)
		if j != nil {
			return []int{tom.index, j.index}
		}
	}
	return nil
}

func binarySearch(nums []*node, target int) *node {
	size := len(nums)
	mid := size / 2
	for {
		mv := nums[mid]
		if mv.value == target {
			return nums[mid]
		}
		if size == 1 {
			return nil
		}
		if mv.value < target {
			return binarySearch(nums[mid:], target)
		} else {
			return binarySearch(nums[:mid], target)
		}
	}

}

//quick sort
func qsort(nums []*node) {
	if len(nums) <= 1 {
		return
	}
	pivot := pivot(len(nums))
	pv := nums[pivot]
	swap(nums, 0, pivot)
	flag := 0
	for i := 1; i < len(nums); i++ {
		if nums[i].value < pv.value {
			flag++
			swap(nums, flag, i)
		}
	}
	swap(nums, 0, flag)
	left := nums[:flag]
	qsort(left)
	right := nums[flag+1:]
	qsort(right)
}

func swap(input []*node, i int, j int) {
	input[i], input[j] = input[j], input[i]
}

func pivot(len int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(len)
}
