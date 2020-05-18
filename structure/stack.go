// stack.go  * Created on  2020/5/14
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

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
