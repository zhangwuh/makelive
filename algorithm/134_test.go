// 134_test.go  * Created on  2020/5/14
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
	"testing"

	"github.com/stretchr/testify/assert"
	"mklive.zhangwuh.com/structure"
)

func TestCheckParenthese(t *testing.T) {
	assert.True(t, checkParenthese("(())()((()())())"))
	assert.False(t, checkParenthese("()())"))
	assert.False(t, checkParenthese("("))
	assert.True(t, checkParenthese("()"))
}

func TestCompleteParenthese(t *testing.T) {
	res, err := completeParenthese("1+2)*3-4)*5-6)))")
	assert.Nil(t, err)
	assert.Equal(t, res, "((1+2)*((3-4)*(5-6)))")

	_, err = completeParenthese("1+2)*3-4)-*5-6)))")
	assert.NotNil(t, err)
}

func TestInfixToPostfix(t *testing.T) {
	res, err := infixToPostfix("a*b/(c-d)+e*(f-g)")
	assert.Nil(t, err)
	assert.Equal(t, res, "ab*cd-/efg-*+")
}

func TestTwoStackQueue(t *testing.T) {
	q := structure.NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, q.Dequeue(), 1)
	assert.Equal(t, q.Dequeue(), 2)
	q.Enqueue(4)
	assert.Equal(t, q.Dequeue(), 3)
	assert.Equal(t, q.Dequeue(), 4)
}
