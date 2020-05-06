// sort.go  * Created on  2020/5/2
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package sort

import (
	"fmt"
	"math/rand"
	"time"
)

func swap(input []int, i, j int) {
	input[i], input[j] = input[j], input[i]
}

type Sorter interface {
	Sort(input []int) []int
}

type bubbleSort struct {
}

func (bs *bubbleSort) Sort(input []int) []int {
	for i := 0; i < len(input); i++ {
		hasSwap := false
		for j := 0; j < len(input)-i-1; j++ {
			left, right := j, j+1
			if input[left] > input[right] {
				swap(input, left, right)
				hasSwap = true
			}
		}
		if !hasSwap {
			fmt.Println(fmt.Sprintf("%s, %d", "no swap,sort is end", i))
			break
		}
	}
	return input
}

type InsertSort struct {
}

func (is *InsertSort) Sort(input []int) []int {
	var output []int
	for _, inserted := range input {
		output = append(output, inserted)
		for i := len(output) - 1; i > 0; i-- {
			if output[i] < output[i-1] {
				swap(output, i, i-1)
			} else {
				fmt.Println(fmt.Sprintf("break on %d", i))
				break
			}
		}
	}
	return output
}

type SelectionSort struct {
}

func (ss *SelectionSort) Sort(input []int) []int {
	for i := 0; i < len(input); i++ {
		var min = i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[min] {
				min = j
			}
		}
		if min != i {
			swap(input, min, i)
		}
	}
	return input
}

type MergeSort struct {
}

func (ss *MergeSort) parallelSort(input []int) chan []int {
	ch := make(chan []int)
	go func() {
		defer close(ch)
		ch <- ss.Sort(input)
	}()
	return ch
}

func (ss *MergeSort) Sort(input []int) []int {
	if len(input) == 1 {
		return input
	}
	m := len(input) / 2
	if m == 0 {
		return input
	}
	lo := ss.parallelSort(input[:m])
	ro := ss.parallelSort(input[m:])
	return ss.merge(<-lo, <-ro)
}

type queue struct {
	data []int
}

func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *queue) first() int {
	return q.data[0]
}

func (q *queue) pop() int {
	defer func() {
		q.data = q.data[1:]
	}()
	return q.first()
}

func (q *queue) append(d int) {
	q.data = append(q.data, d)
}

func (ss *MergeSort) merge(lo []int, ro []int) (out []int) {
	lq := &queue{lo}
	rq := &queue{ro}

	for !lq.isEmpty() || !rq.isEmpty() {
		if lq.isEmpty() {
			out = append(out, rq.pop())
		} else if rq.isEmpty() {
			out = append(out, lq.pop())
		} else {
			if lq.first() < rq.first() {
				out = append(out, lq.pop())
			} else {
				out = append(out, rq.pop())
			}
		}
	}
	return
}

type QuickSort struct {
	count int
}

func (ss *QuickSort) parallelSort(input []int) chan []int {
	ch := make(chan []int)
	go func() {
		defer close(ch)
		ch <- ss.Sort(input)
	}()
	return ch
}

func (qs *QuickSort) Sort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	pivot := r.Intn(len(input))

	pv := input[pivot]
	lq := input[:pivot]
	rq := input[pivot+1:]

	slq := &queue{lq}
	srq := &queue{rq}
	var ol, or []int
	for {
		if slq.isEmpty() {
			break
		}
		qs.count++
		v := slq.pop()
		if v >= pv {
			or = append(or, v)
		} else {
			ol = append(ol, v)
		}
	}
	for {
		if srq.isEmpty() {
			break
		}
		qs.count++
		v := srq.pop()
		if v >= pv {
			or = append(or, v)
		} else {
			ol = append(ol, v)
		}
	}

	return append(append(qs.Sort(ol), pv), qs.Sort(or)...)
}
