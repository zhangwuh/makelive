// 134.go  * Created on  2020/5/14
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package algorithm

import (
	"fmt"

	"mklive.zhangwuh.com/structure"
)

func checkParenthese(input string) bool { //(())()((()())()) true,()()) false
	stack := structure.NewLinkedStack(nil)
	for _, c := range input {
		if c == ')' {
			n := stack.Pop()
			if n != '(' {
				return false
			}
		} else {
			stack.Push(c)
		}
	}
	return stack.IsEmpty()
}

var symbos = []rune{'+', '-', '*', '/'}

func runesContains(runes []rune, r rune) bool {
	for _, v := range runes {
		if v == r {
			return true
		}
	}
	return false
}

func completeParenthese(input string) (string, error) { //1+2)*3-4)*5-6))) -> ((1+2)*((3-4)*(5-6)))
	stack := structure.NewLinkedStack(nil)
	symbolStack := structure.NewLinkedStack(nil)
	for _, c := range input {
		if runesContains(symbos, c) {
			symbolStack.Push(string(c))
		} else if c == ')' {
			tom, jerry := stack.Pop(), stack.Pop()
			symbol := symbolStack.Pop()
			if tom == nil || jerry == nil || symbol == nil {
				return "", fmt.Errorf("invalid input")
			}
			stack.Push(fmt.Sprintf("(%s%s%s)", jerry.(string), symbol.(string), tom.(string)))
		} else {
			stack.Push(string(c))
		}
	}
	if !symbolStack.IsEmpty() {
		return "", fmt.Errorf("invalid input")
	}
	return stack.Pop().(string), nil
}

func infixToPostfix(input string) (string, error) { //2*3/(2-1)+3*(4-1) -> 23*21-/341-*+
	stack := structure.NewLinkedStack(nil)
	symbolStack := structure.NewLinkedStack(nil)
	for _, c := range input {
		if c == '(' {
			continue
		}
		if runesContains(symbos, c) {
			symbolStack.Push(string(c))
		} else if c == ')' {
			tom, jerry := stack.Pop(), stack.Pop()
			symbol := symbolStack.Pop()
			if tom == nil || jerry == nil || symbol == nil {
				return "", fmt.Errorf("invalid input")
			}
			stack.Push(fmt.Sprintf("%s%s%s", jerry.(string), tom.(string), symbol.(string)))
		} else {
			stack.Push(string(c))
		}
	}
	ns := stack.Inverse()
	nss := symbolStack.Inverse()
	for !ns.IsEmpty() {
		tom, jerry := ns.Pop(), ns.Pop()
		if jerry == nil {
			jerry = ""
		}
		symbol := nss.Pop()
		if symbol == nil {
			symbol = ""
		}
		ns.Push(fmt.Sprintf("%s%s%s", tom.(string), jerry.(string), symbol.(string)))
	}
	return ns.Peek().(string), nil
}
