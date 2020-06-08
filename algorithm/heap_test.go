package algorithm

import (
	"fmt"
	"testing"
)

func TestNewIndexMinPQ(t *testing.T) {
	q := NewIndexMinPQ([]int{5, 9, 10, 2, 6, 1, 0})
	q.Insert(7, 5)
	q.Insert(8, 11)
	q.Print()
	fmt.Println("---------------")
	q.DelRoot()
	q.Print()
	fmt.Println("---------------")
	q.DelRoot()
	q.Print()
}
