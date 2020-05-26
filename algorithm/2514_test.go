// 2514_test.go  * Created on  2020/5/26
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

func TestSortDomains(t *testing.T) {
	ds := []*Domain{
		NewDomain("ee.princeton.edu"),
		NewDomain("cs.princeton.edu"),
		NewDomain("princeton.edu"),
		NewDomain("cnn.com"),
		NewDomain("google.com"),
		NewDomain("apple.com"),
		NewDomain("www.cs.princeton.edu"),
		NewDomain("bolle.cs.princeton.edu"),
	}
	output := SortDomains(ds)
	for _, v := range output {
		fmt.Println(v.(*Domain).url)
	}
}
