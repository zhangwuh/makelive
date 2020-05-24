// heap.go  * Created on  2020/5/24
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package algorithm

import "fmt"

type IndexPQ interface {
	Insert(index, value int)
	swim(i int)
	sink(i int)
}

type IndexMinPQ struct {
	pq     []int
	source []int
}

func (q *IndexMinPQ) swim(i int) {
	for i/2 > 0 {
		parent := i / 2
		if q.value(i) < q.value(parent) {
			q.swap(i, parent)
			i = parent
		} else {
			return
		}
	}
}

func (q *IndexMinPQ) index(i int) int {
	return q.pq[i]
}

func (q *IndexMinPQ) value(i int) int {
	if i == 0 {
		panic(fmt.Errorf("invalid index"))
	}
	return q.source[q.pq[i]]
}

func (q *IndexMinPQ) min(i, j int) int {
	if q.value(i) <= q.value(j) {
		return i
	}
	return j
}
func (q *IndexMinPQ) sink(i int) {
	for 2*i < len(q.pq) {
		lc := 2 * i
		var min int
		if lc+1 < len(q.pq) {
			lr := lc + 1
			min = q.min(lc, lr)
		} else {
			min = lc
		}
		if q.value(min) < q.value(i) {
			q.swap(i, min)
			i = min
		} else {
			return
		}
	}
}

func (q *IndexMinPQ) swap(i, j int) {
	q.pq[i], q.pq[j] = q.pq[j], q.pq[i]
}

func NewIndexMinPQ(source []int) *IndexMinPQ {
	pq := &IndexMinPQ{
		pq: []int{0},
	}
	for i, v := range source {
		pq.Insert(i, v)
	}
	return pq
}

func (pq *IndexMinPQ) Print() {
	for _, i := range pq.pq[1:] {
		fmt.Println(pq.source[i])
	}
}

func (pq *IndexMinPQ) Insert(i, val int) {
	pq.source = append(pq.source, val)
	pq.pq = append(pq.pq, i)
	index := len(pq.pq) - 1
	pq.swim(index)
}

func (pq *IndexMinPQ) Min() int {
	return pq.value(1)
}

func (pq *IndexMinPQ) DelMin() int {
	min := pq.Min()
	pq.swap(1, len(pq.pq)-1)
	pq.pq = pq.pq[0 : len(pq.pq)-1]
	pq.sink(1)
	return min
}
