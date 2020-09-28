package inorderTraversal

import "testing"

func NewTreeNode(t int) *TreeNode  {
	return &TreeNode{Val: t}
}
func TestInorderTraversal(t *testing.T) {
	tr := NewTreeNode(1)
	tr.Right = NewTreeNode(9)
	tr.Right = NewTreeNode(5)
	tr.Right = NewTreeNode(6)
	tr.Left = NewTreeNode(5)
	tr.Left = NewTreeNode(0)

	info := inorderTraversal(tr)
	t.Log(info)
}
