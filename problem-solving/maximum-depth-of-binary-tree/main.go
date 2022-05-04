package maximum_depth_of_binary_tree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxDepthHelper(root, 0)
}

func maxDepthHelper(root *TreeNode, count int) int {
	if root == nil {
		return count
	}
	left := maxDepthHelper(root.Left, count+1)

	right := maxDepthHelper(root.Right, count+1)

	if left < right {
		return right
	} else {
		return left
	}

}
