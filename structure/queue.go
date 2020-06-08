package structure

type Queue interface {
	Dequeue() interface{}
	Enqueue(interface{})
}

type TwoStackQueue struct {
	s1 *LinkedStack //enq
	s2 *LinkedStack //deq only
}

func NewQueue() *TwoStackQueue {
	return &TwoStackQueue{
		s1: NewLinkedStack(nil),
		s2: NewLinkedStack(nil),
	}
}

func (ts *TwoStackQueue) Dequeue() interface{} {
	if !ts.s1.IsEmpty() {
		ts.s1.PushAll(ts.s2)
		return ts.s2.Pop()
	}
	return ts.s2.Pop()
}

func (ts *TwoStackQueue) Enqueue(e interface{}) {
	if ts.s1.IsEmpty() {
		ts.s2.PushAll(ts.s1)
	}
	ts.s1.Push(e)
}
