// load_balance_test.go  * Created on  2020/5/26
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package algorithm

import "testing"

func TestBalancer_Take(t *testing.T) {
	balancer := NewBalancer(3)
	tasks := []*task{
		{"a", 10},
		{"b", 5},
		{"c", 20},
		{"d", 15},
		{"e", 11},
	}
	balancer.Take(tasks)
}
