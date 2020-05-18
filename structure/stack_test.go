// stack_test.go  * Created on  2020/5/14
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedStack(t *testing.T) {
	stack := NewLinkedStack(&node{Value: "1"})

	stack.Push(&node{Value: "2"})
	stack.Push(&node{Value: "3"})

	assert.Equal(t, stack.Pop(), "3")
	assert.Equal(t, stack.Pop(), "2")
	assert.Equal(t, stack.Pop(), "1")
	assert.True(t, stack.IsEmpty())
	assert.Nil(t, stack.Pop())
}
