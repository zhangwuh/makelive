package algorithm

import (
	"fmt"
	"testing"
)

func Test_indexCounterSorter(t *testing.T) {
	nodes := []node{
		{"a", 1},
		{"b", 4},
		{"c", 3},
		{"d", 1},
		{"e", 2},
		{"f", 2},
		{"g", 0},
		{"h", 1},
		{"i", 4},
		{"j", 3},
	}

	sorter := newIndexCounterSorter(nodes, 5)
	fmt.Println(sorter.sort())
}

func Test_LSD(t *testing.T) {
	input := []string{"abce", "dcbb", "xahb", "cccc"}
	for j := 0; j < 4; j++ {
		ns := []node{}
		for _, i := range input {
			ns = append(ns, node{i, int(i[j])})
		}
		sorter := newIndexCounterSorter(ns, 256)
		for n, m := range sorter.sort() {
			input[n] = m.name
		}
	}
	fmt.Println(input)
}
