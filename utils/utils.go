// utils.go  * Created on  2020/5/27
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package utils

import (
	"math/rand"
	"time"
)

type DeQue interface {
	Add(e interface{})
	Offer() interface{}
	Stack
}

type Stack interface {
	Pop() interface{}
	Push(e interface{})
	Size() int
	IsEmpty() bool
}

type deque struct {
	stack
	stackMode bool
}

func (dq *deque) inverse() {
	dq.stackMode = !dq.stackMode
}

func (dq *deque) next() interface{} {
	if dq.IsEmpty() {
		return nil
	}
	if dq.stackMode {
		return dq.Pop()
	}
	return dq.Offer()
}

func (dq *deque) putback(e interface{}) {
	if dq.stackMode {
		dq.Push(e)
	} else {
		dq.Add(e)
	}
}

func (dq *deque) Add(e interface{}) {
	var newarr []interface{}
	newarr = append(newarr, e)
	newarr = append(newarr, dq.arr...)
	dq.arr = newarr
}

func (dq *deque) Offer() interface{} {
	if dq == nil || dq.stack.IsEmpty() {
		return nil
	}
	defer func() {
		dq.arr = dq.arr[1:]
	}()
	return dq.arr[0]
}

type stack struct {
	arr []interface{}
}

func (rs *stack) IsEmpty() bool {
	return rs == nil || rs.Size() == 0
}

func (rs *stack) Pop() interface{} {
	if rs == nil || rs.IsEmpty() {
		return nil
	}
	defer func() {
		rs.arr = rs.arr[0 : rs.Size()-1]
	}()
	return rs.arr[rs.Size()-1]
}

func (rs *stack) Push(r interface{}) {
	rs.arr = append(rs.arr, r)
}

func (rs *stack) Size() int {
	return len(rs.arr)
}

type NumberSorter struct {
	Desc bool
}

func pivot(len int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(len)
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

func ContainsInt(slice []int, n int) bool {
	for _, v := range slice {
		if v == n {
			return true
		}
	}
	return false
}
