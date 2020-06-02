// bloom_filter_test.go  * Created on  2020/6/2
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package practise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloomFilter(t *testing.T) {
	filter := NewBloomFilter(0.001, 1000)
	filter.Set([]byte("abcd"))
	filter.Set([]byte("aaaa"))
	filter.Set([]byte("bbdcd"))

	assert.Equal(t, filter.Get([]byte("abcd")), true)
	assert.Equal(t, filter.Get([]byte("bbdcd")), true)

	assert.Equal(t, filter.Get([]byte("abc")), false)
}
