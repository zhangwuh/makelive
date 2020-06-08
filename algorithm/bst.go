package algorithm

//Binary search tree
type TreeNode struct {
	key   string
	value interface{}
	left  *TreeNode
	right *TreeNode
}

type Table interface {
	Get(key string) interface{}
	Put(key string, val interface{})
}

func NewTreeNode(key, val string) *TreeNode {
	return &TreeNode{key: key, value: val}
}

func (n *TreeNode) Key() string {
	return n.key
}

func (n *TreeNode) Value() interface{} {
	return n.value
}

func (n *TreeNode) compareTo(to *TreeNode) int {
	if n.key == to.key {
		return 0
	}
	if n.key > to.key {
		return 1
	}
	return -1
}

type BST struct {
	root *TreeNode
}

func (bst *BST) leftTree() *BST {
	if bst.root == nil || bst.root.left == nil {
		return nil
	}
	return &BST{bst.root.left}
}

func (bst *BST) rightTree() *BST {
	if bst.root == nil || bst.root.right == nil {
		return nil
	}
	return &BST{bst.root.right}
}

func (bst *BST) Get(key string) interface{} {
	if bst == nil || bst.root == nil {
		return nil
	}

	if bst.root.key == key {
		return bst.root.value
	}

	if key < bst.root.key {
		return bst.leftTree().Get(key)
	}
	return bst.rightTree().Get(key)
}

func (bst *BST) Put(key string, value interface{}) {
	node := &TreeNode{key: key, value: value}
	if bst.root == nil || bst.root.compareTo(node) == 0 {
		bst.root = node
		return
	}
	if node.compareTo(bst.root) < 0 {
		if bst.leftTree() == nil {
			bst.root.left = node
		} else {
			bst.leftTree().Put(key, value)
		}
	} else {
		if bst.rightTree() == nil {
			bst.root.right = node
		} else {
			bst.rightTree().Put(key, value)
		}
	}
}
