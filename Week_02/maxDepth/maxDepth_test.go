package maxDepth

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		return int(math.Max(float64(left), float64(right)) + 1)
	}
}
