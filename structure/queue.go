// queue.go  * Created on  2020/5/16
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

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
