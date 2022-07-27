package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var list []*TreeNode
	var preorder func(root *TreeNode)

	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		list = append(list, root)

		if root.Left != nil {
			preorder(root.Left)
		}
		if root.Right != nil {
			preorder(root.Right)
		}
	}

	preorder(root)

	for i := 0; i < len(list); i++ {
		list[i].Left = nil
		if i == len(list)-1 {
			list[i].Right = nil
		} else {
			list[i].Right = list[i+1]
		}
	}
}
