// 4_test.go  * Created on  2020/5/26
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

func Test_Merge(t *testing.T) {
	n1 := []int{1, 2}
	n2 := []int{3, 4}
	assert.Equal(t, findMedianSortedArrays(n1, n2), 2.5)

	n1 = []int{1, 3, 5, 7, 9, 11}
	n2 = []int{0, 2, 4, 6, 8}

	assert.Equal(t, findMedianSortedArrays(n1, n2), float64(5))

	n1 = []int{0, 0}
	n2 = []int{0, 0}

	assert.Equal(t, findMedianSortedArrays(n1, n2), float64(0))
}
