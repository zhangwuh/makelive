// q1_test.go  * Created on  2020/5/26
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

func TestQ1(t *testing.T) {
	input := []int{7, 2, 3, 5, 4, 15, 11}
	r := twoSumOnMap(input, 11)
	assert.Equal(t, input[r[0]]+input[r[1]], 11)

}
