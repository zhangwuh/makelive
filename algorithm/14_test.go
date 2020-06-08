package algorithm

import (
	"fmt"
	"testing"
)

func TestPickEqualNumbers(t *testing.T) {
	input := []int{0, 3, 0, 4, 6, 9, 4, 6}
	fmt.Println(pickEqualNumbers(input))
}

func TestPickCommons(t *testing.T) {
	input1 := []int{0, 3, 6, 9, 10, 12, 19, 21}
	input2 := []int{0, 3, 5, 9, 12, 19, 25, 29}
	for _, v := range PickCommons(input1, input2) {
		fmt.Println(v)
	}
}
