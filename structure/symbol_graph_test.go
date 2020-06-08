// symbol_graph_test.go  * Created on  2020/6/7
// Copyright (c) 2020 YueTu
// YueTu TECHNOLOGY CO.,LTD. All Rights Reserved.
//
// This software is the confidential and proprietary information of
// YueTu Ltd. ("Confidential Information").
// You shall not disclose such Confidential Information and shall use
// it only in accordance with the terms of the license agreement you
// entered into with YueTu Ltd.

package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolGraph(t *testing.T) {
	sg := NewSymbolGraph([]string{
		"BJS NYC", "CTU BJS", "BKK NYC",
	})
	assert.Equal(t, len(sg.Adj("BJS")), 2)
	assert.Equal(t, len(sg.Adj("BKK")), 1)
}
