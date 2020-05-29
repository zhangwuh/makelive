package leetcode

// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

type ListNode struct {
	Val  int
	Next *ListNode
}

var zeroNode = ListNode{
	Val: -1,
}

func isZero(n ListNode) bool {
	return n.Val == -1
}

func Pop(l *ListNode) ListNode {
	defer func() {
		if l == nil {
			return
		}
		if l.Next == nil {
			l.Val = -1
			return
		}
		next := l.Next
		l.Next = next.Next
		l.Val = next.Val
	}()
	return *l
}

func getVal(l ListNode) int {
	if isZero(l) {
		return 0
	}
	return l.Val
}

func Push(head *ListNode, v int) {
	old := &ListNode{
		Val:  head.Val,
		Next: head.Next,
	}
	head.Val = v
	head.Next = old
}

//O(max(m, n))

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var last *ListNode
	var head *ListNode
	l3 := &ListNode{}
	for {
		n1 := Pop(l1)
		n2 := Pop(l2)
		n3 := Pop(l3)
		if isZero(n1) && isZero(n2) && isZero(n3) {
			break
		}
		v := getVal(n1) + getVal(n2) + getVal(n3)
		if v >= 10 {
			Push(l3, 1)
			v = v - 10
		}
		if head == nil {
			head = &ListNode{
				Val: v,
			}
			last = head
		} else {
			current := &ListNode{Val: v}
			last.Next = current
			last = current
		}

	}
	return head
}
