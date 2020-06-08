package structure

type Stack interface {
	Pop() interface{}
	Push(interface{})
	IsEmpty() bool
	Peek() interface{}
}

type node struct {
	Value interface{}
	next  *node
}

type LinkedStack struct {
	head *node
}

func NewLinkedStack(head *node) *LinkedStack {
	return &LinkedStack{
		head: head,
	}
}
func (ls *LinkedStack) IsEmpty() bool {
	return ls.head == nil
}

func (ls *LinkedStack) Peek() interface{} {
	return ls.head.Value
}

func (ls *LinkedStack) Pop() interface{} {
	if ls.IsEmpty() {
		return nil
	}
	defer func() {
		ls.head = ls.head.next
	}()
	return ls.head.Value
}

func (ls *LinkedStack) Push(v interface{}) {
	n := &node{
		Value: v,
	}
	n.next = ls.head
	ls.head = n
}

func (ls *LinkedStack) Inverse() *LinkedStack {
	ns := &LinkedStack{nil}
	for !ls.IsEmpty() {
		ns.Push(ls.Pop())
	}
	return ns
}

func (ls *LinkedStack) PushAll(ts *LinkedStack) {
	for !ls.IsEmpty() {
		ts.Push(ls.Pop())
	}
}
