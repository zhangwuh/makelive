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

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	source := strconv.Itoa(x)
	isNegative := false
	if x < 0 {
		isNegative = true
		source = source[1:]
	}
	var out []rune
	for i := len(source) - 1;i >= 0 ;i-- {
		out = append(out,  rune(source[i]))
	}
	if overflow(out, isNegative) {
		return 0
	}
	res, _ := strconv.Atoi(string(out))
	if isNegative {
		return -res
	}
	return res
}

var maxRunes = strconv.Itoa(math.MaxInt32)
var minRunes = strconv.Itoa(math.MinInt32)[1:]
var maxLen = len(maxRunes)

func overflow(source []rune, isNegative bool) bool {
	if len(source) < maxLen {
		return false
	}

	if len(source) == maxLen {
		if isNegative {
			return string(source) > maxRunes
		} else {
			return string(source) > minRunes
		}
	} else {
		return true
	}
}
