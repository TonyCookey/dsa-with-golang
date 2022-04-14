package searchBST

// TreeNode Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := searchBST(root.Left, val)
	if left != nil {
		return left
	}
	right := searchBST(root.Right, val)
	if right != nil {
		return right
	}
	return nil
}

func searchBST2(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST2(root.Left, val)
	} else {
		return searchBST2(root.Right, val)
	}
}
