package construct_binary_tree_frmo_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}

	mid := findIndex(inorder, preorder[0])
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])

	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}

func findIndex(xi []int, x int) int {
	for i, v := range xi {
		if v == x {
			return i
		}
	}
	return 0
}
