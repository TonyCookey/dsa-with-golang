package balanced_binary_tree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	balanced, _ := isBalancedHelper(root)

	return balanced
}

func isBalancedHelper(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftBalanced, leftHeight := isBalancedHelper(root.Left)
	rightBalanced, rightHeight := isBalancedHelper(root.Right)

	balanced := abs(leftHeight-rightHeight) <= 1 && leftBalanced && rightBalanced

	return balanced, max(leftHeight, rightHeight) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(diff int) int {

	if diff > 0 {
		return diff
	}
	return -diff
}
