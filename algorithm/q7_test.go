// q7_test.go  * Created on  2020/6/2
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
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))
	assert.Equal(t, 0, reverse(1323423233))
	assert.Equal(t, 0, reverse(-1323423233))
	assert.Equal(t, 2147483641, reverse(1463847412))
	assert.Equal(t, -2147483641, reverse(-1463847412))
}
