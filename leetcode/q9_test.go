// q9_test.go  * Created on  2020/6/3
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	assert.True(t, isPalindrome(121))
	assert.False(t, isPalindrome(-121))
	assert.False(t, isPalindrome(10))
}
