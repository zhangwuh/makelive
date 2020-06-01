// q6_test.go  * Created on  2020/6/1
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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convert(t *testing.T) {
	assert.Equal(t, "LCIRETOESIIGEDHN", convert("LEETCODEISHIRING", 3))
	assert.Equal(t, "LDREOEIIECIHNTSG", convert("LEETCODEISHIRING", 4))
	assert.Equal(t, "AB", convert("AB", 1))
	fmt.Println(0 % 1)
}
