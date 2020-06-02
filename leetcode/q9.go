// q9.go  * Created on  2020/6/3
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package leetcode

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	var reverse int
	copy := x
	for copy != 0 {
		num := copy % 10
		if reverse > 214748364 || reverse == 214748364 && num > 7 { //overflow
			return false
		}
		reverse = reverse*10 + num
		copy /= 10
	}
	return reverse == x
}
