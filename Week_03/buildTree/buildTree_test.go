package buildTree

import "testing"

func TestBuildTree(t *testing.T) {
	// 前序遍历
	preorder := []int{3, 9, 20, 15, 7}
	// 中序遍历
	inorder := []int{9, 3, 15, 20, 7}
	info := buildTree(preorder, inorder)
	t.Log(info)

}
