// q5_test.go  * Created on  2020/5/30
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

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, longestPalindrome("babad"), "bab")
	assert.Equal(t, longestPalindrome("cbbd"), "bb")
	assert.Equal(t, longestPalindrome("a"), "a")
	assert.Equal(t, longestPalindrome(""), "")
	assert.Equal(t, longestPalindrome("abb"), "bb")
}
