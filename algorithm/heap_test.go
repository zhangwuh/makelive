// heap_test.go  * Created on  2020/5/24
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

func TestNewIndexMinPQ(t *testing.T) {
	q := NewIndexMinPQ([]int{5, 9, 10, 2, 6, 1, 0})
	q.Insert(7, 5)
	q.Insert(8, 11)
	q.Print()
	fmt.Println("---------------")
	q.DelMin()
	q.Print()
	fmt.Println("---------------")
	q.DelMin()
	q.Print()
}
