package leetcode

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	s1 := &ListNode{
		Val: 2,
	}
	Push(s1, 4)
	Push(s1, 3)
	s2 := &ListNode{
		Val: 5,
	}
	Push(s2, 6)
	Push(s2, 4)
	res := addTwoNumbers(s1, s2)
	for {
		fmt.Println(res.Val)
		if res.Next == nil {
			break
		} else {
			res = res.Next
		}
	}
}
