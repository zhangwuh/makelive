// bst_test.go  * Created on  2020/5/26
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

func buildTree() *BST {
	root := &TreeNode{key: "S", value: "s"}
	root.left = &TreeNode{key: "E", value: "e"}
	root.right = &TreeNode{key: "X", value: "x"}

	root.left.left = &TreeNode{key: "A", value: "a"}
	root.left.right = &TreeNode{key: "R", value: "r"}

	root.left.left.right = &TreeNode{key: "C", value: "c"}
	root.left.right.left = &TreeNode{key: "H", value: "h"}
	root.left.right.left.right = &TreeNode{key: "M", value: "m"}
	return &BST{
		root: root,
	}
}

func TestBST_Get(t *testing.T) {
	bst := buildTree()

	assert.Equal(t, bst.Get("R"), "r")
	assert.Equal(t, bst.Get("Z"), nil)
}

func TestBST_Put(t *testing.T) {
	bst := &BST{}
	bst.Put("S", "s")
	bst.Put("E", "s")
	bst.Put("A", "s")
	bst.Put("C", "s")
	bst.Put("R", "s")
	bst.Put("H", "s")
	bst.Put("M", "s")
	bst.Put("X", "s")

	assert.Equal(t, bst.root.right.key, "X")
}
