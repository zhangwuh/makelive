// q3.go  * Created on  2020/5/29
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度
package leetcode

func lengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubRunes([]rune(s))
}

func lengthOfLongestSubRunes(runes []rune) int {
	if len(runes) == 0 {
		return 0
	}
	index := make(map[rune]int)
	var max, nmax int
	var i int
	for {
		if i == len(runes) {
			break
		}
		r := runes[i]
		if last, ok := index[r]; ok {
			i = last + 1
			if max < nmax {
				max = nmax
			}
			nmax = 0
			index = make(map[rune]int)
			continue
		}
		nmax++
		index[r] = i
		i++
	}
	if max < nmax {
		max = nmax
	}
	return max
}
