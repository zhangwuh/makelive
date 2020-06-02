// q7.go  * Created on  2020/6/2
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package algorithm

const (
	ceil  = 214748364
	floor = -214748364
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var res int
	for x != 0 {
		low := x % 10
		if res > ceil || (res == ceil && low > 7) {
			return 0
		}
		if res < floor || (res == floor && low < -8) {
			return 0
		}
		x /= 10
		res = res*10 + low
	}
	return res
}
