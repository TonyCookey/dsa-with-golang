package is_symmetric_tree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricRecursive(root.Left, root.Right)
}

func isSymmetricRecursive(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricRecursive(left.Left, right.Right) && isSymmetricRecursive(left.Right, right.Left)

}
