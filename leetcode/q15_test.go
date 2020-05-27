// q15_test.go  * Created on  2020/5/27
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
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	for _, v := range res {
		fmt.Println(v)
	}

	nums = []int{0, 0, 0}
	res = threeSum(nums)
	for _, v := range res {
		fmt.Println(v)
	}
}
