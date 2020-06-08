package sort

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	input := []int{3, 2, 1}

	swap(input, 1, 2)
	fmt.Println(input)
}

func TestBubbleSort_Sort(t *testing.T) {
	input := []int{3, 2, 1, 10, 20, 5, 1, 0}
	st := &bubbleSort{}
	fmt.Println(st.Sort(input))

	input = []int{1, 4, 3, 5, 6, 7}
	fmt.Println(st.Sort(input))
}

func TestInsertSort_Sort(t *testing.T) {
	input := []int{3, 2, 1, 10, 20, 5, 1, 0}
	st := &InsertSort{}
	fmt.Println(st.Sort(input))
	input = []int{1, 4, 3, 5, 6, 7}
	fmt.Println(st.Sort(input))
}

func TestSelectionSort_Sort(t *testing.T) {
	input := []int{3, 2, 1, 10, 20, 5, 1, 0}
	st := &SelectionSort{}
	fmt.Println(st.Sort(input))
	input = []int{1, 4, 3, 5, 6, 7}
	fmt.Println(st.Sort(input))
}

func TestMergeSort_Sort(t *testing.T) {
	input := []int{3, 2, 1, 10, 20, 5, 1, 0}
	st := &MergeSort{}
	fmt.Println(st.Sort(input))
}

func TestPop(t *testing.T) {
	input := &queue{[]int{3, 2, 1, 10, 20, 5, 1, 0}}
	for !input.isEmpty() {
		fmt.Println(input.peek())
	}
}

func TestQuickSort_Sort(t *testing.T) {
	input := []int{3, 2, 1, 10, 20, 5, 1, 9}
	st := &QuickSort{}
	fmt.Println(st.Sort(input))
	fmt.Println(st.count)

	input = []int{10, 9, 8, 7, 6, 5, 4, 3}
	fmt.Println(st.Sort(input))
	fmt.Println(st.count)
}

func TestHillSort_Sort(t *testing.T) {
	input := []int{3, 2, 1, 10, 20, 5, 1, 9}
	st := &HillSort{}
	fmt.Println(st.Sort(input))
}

type stringCompare string

func (sc stringCompare) Compare(to Comparable) int {
	if string(sc) == string(to.(stringCompare)) {
		return 0
	}
	if string(sc) >= string(to.(stringCompare)) {
		return 1
	}
	return 1
}

func TestNewQuickSort_Sort(t *testing.T) {
	qs := &NewQuickSort{Desc: true}
	var input []Comparable
	for _, v := range []stringCompare{"B", "C", "A", "G", "F", "E"} {
		input = append(input, v)
	}
	qs.Sort(input)
	for _, v := range input {
		println(string(v.(stringCompare)))
	}
}
