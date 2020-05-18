// 148.go  * Created on  2020/5/17
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package algorithm

import (
	"mklive.zhangwuh.com/sort"
	"mklive.zhangwuh.com/structure"
)

func pickEqualNumbers(input []int) []int {
	sorted := (&sort.QuickSort{}).Sort(input)
	var current int
	var output []int
	for i := 0; i < len(sorted); i++ {
		if i == 0 {
			current = sorted[0]
			continue
		}
		if sorted[i] == current {
			output = append(output, current)
		} else {
			current = sorted[i]
		}
	}
	return output
}

func PickCommons(i1 []int, i2 []int) []int {
	s1 := structure.NewQueue()
	s2 := structure.NewQueue()
	for _, v := range i1 {
		s1.Enqueue(v)
	}
	for _, v := range i2 {
		s2.Enqueue(v)
	}

	var out []int
	v := s1.Dequeue()
	c := s2.Dequeue()
	for {
		if v == nil || c == nil {
			break
		}
		if v == c {
			out = append(out, v.(int))
			v = s1.Dequeue()
			c = s2.Dequeue()
			continue
		}

		if v.(int) > c.(int) {
			c = s2.Dequeue()
			continue
		} else {
			v = s1.Dequeue()
			continue
		}

	}
	return out
}
