// utils.go  * Created on  2020/5/27
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package leetcode

type NumberSorter struct {
	Desc bool
}

func (qs *NumberSorter) Sort(input []int) {
	if len(input) <= 1 {
		return
	}
	pivot := pivot(len(input))
	qs.doSort(input, pivot)
}

func (qs *NumberSorter) swap(input []int, i int, j int) {
	input[i], input[j] = input[j], input[i]
}

func (qs *NumberSorter) doSort(input []int, pivot int) {
	pv := input[pivot]
	qs.swap(input, pivot, 0)
	flag := 0
	for i := 1; i < len(input); i++ {
		compare := input[i] < pv
		if qs.Desc {
			compare = !compare
		}
		if compare {
			flag++
			qs.swap(input, i, flag)
		}
	}
	qs.swap(input, 0, flag)
	left := input[:flag]
	right := input[flag+1:]
	qs.Sort(left)
	qs.Sort(right)
}
