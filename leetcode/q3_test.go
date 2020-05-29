// q3_test.go  * Created on  2020/5/29
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
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		input string
		res   int
	}{
		{
			"abcabcbb", 3,
		},
		{
			"bbbbb", 1,
		},
		{
			"pwwkew", 3,
		},
		{
			"dvdf", 3,
		},
		{
			"ohomm", 3,
		},
	}
	for _, c := range cases {
		m := lengthOfLongestSubstring(c.input)
		fmt.Println(m)
		assert.Equal(t, m, c.res)
	}

}

func Benchmark_lengthOfLongestSubstring(t *testing.B) {
	sample, _ := ioutil.ReadFile("q3")
	for i := 0; i < t.N; i++ {
		fmt.Println(lengthOfLongestSubstring(string(sample)))
	}
}
