package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return DFSHelper(root, 0, targetSum)
}

func DFSHelper(root *TreeNode, sum, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val+sum == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	sum += root.Val

	left := DFSHelper(root.Left, sum, targetSum)
	if left {
		return left
	}

	right := DFSHelper(root.Right, sum, targetSum)
	if right {
		return right
	}
	return false

}
func DFSHelperCleaner(root *TreeNode, sum, targetSum int) bool {
	if root == nil {
		return false
	}

	sum += root.Val

	if sum == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return DFSHelper(root.Left, sum, targetSum) || DFSHelper(root.Right, sum, targetSum)
}
