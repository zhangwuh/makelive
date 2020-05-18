// 14_Test.go  * Created on  2020/5/17
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
	"testing"
)

func TestPickEqualNumbers(t *testing.T) {
	input := []int{0, 3, 0, 4, 6, 9, 4, 6}
	fmt.Println(pickEqualNumbers(input))
}

func TestPickCommons(t *testing.T) {
	input1 := []int{0, 3, 6, 9, 10, 12, 19, 21}
	input2 := []int{0, 3, 5, 9, 12, 19, 25, 29}
	for _, v := range PickCommons(input1, input2) {
		fmt.Println(v)
	}
}
