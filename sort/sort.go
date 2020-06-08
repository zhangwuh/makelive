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

func (q *queue) peek() int {
	return q.data[0]
}

func (q *queue) take() int {
	defer func() {
		q.data = q.data[1:]
	}()
	return q.peek()
}

func (q *queue) append(d int) {
	q.data = append(q.data, d)
}

func (ss *MergeSort) merge(lo []int, ro []int) (out []int) {
	lq := &queue{lo}
	rq := &queue{ro}

	for !lq.isEmpty() || !rq.isEmpty() {
		if lq.isEmpty() {
			out = append(out, rq.take())
		} else if rq.isEmpty() {
			out = append(out, lq.take())
		} else {
			if lq.peek() < rq.peek() {
				out = append(out, lq.take())
			} else {
				out = append(out, rq.take())
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

	pivot := pivot(len(input))

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
		v := slq.take()
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
		v := srq.take()
		if v >= pv {
			or = append(or, v)
		} else {
			ol = append(ol, v)
		}
	}

	return append(append(qs.Sort(ol), pv), qs.Sort(or)...)
}

func pivot(len int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(len)
}

type HillSort struct{}

func (qs *HillSort) Sort(input []int) []int {
	for step := len(input) / 2; step > 0; step = step / 2 {
		for i := 0; i < len(input)/step; i++ {
			for j := i; j+step < len(input); j += step {
				if input[j] > input[j+step] {
					input[j], input[j+step] = input[j+step], input[j]
				}
			}
		}
	}
	return input
}

type Comparable interface {
	Compare(to Comparable) int
}

type NewQuickSort struct {
	Desc bool
}

func (qs *NewQuickSort) Sort(input []Comparable) {
	if len(input) <= 1 {
		return
	}
	pivot := pivot(len(input))
	qs.doSort(input, pivot)
}

func (qs *NewQuickSort) swap(input []Comparable, i int, j int) {
	input[i], input[j] = input[j], input[i]
}

func (qs *NewQuickSort) doSort(input []Comparable, pivot int) {
	pv := input[pivot]
	qs.swap(input, pivot, 0)
	flag := 0
	for i := 1; i < len(input); i++ {
		compare := input[i].Compare(pv)
		if qs.Desc {
			compare = -compare
		}
		if compare <= 0 {
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
