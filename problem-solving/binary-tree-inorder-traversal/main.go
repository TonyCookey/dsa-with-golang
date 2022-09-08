package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var list []int

	var DFSInOrder func(node *TreeNode)

	DFSInOrder = func(node *TreeNode) {
		if node.Left != nil {
			DFSInOrder(node.Left)
		}

		list = append(list, node.Val)

		if node.Right != nil {
			DFSInOrder(node.Right)
		}
	}
	DFSInOrder(root)

	return list
}
