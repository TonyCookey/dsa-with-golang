package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	return min(leftDepth, rightDepth) + 1
}

func min(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}

	if x < y {
		return x
	}
	return y
}
