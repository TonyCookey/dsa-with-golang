package count_good_nodes_in_binary_tree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count int

	var DFS func(root *TreeNode, max_num int)

	DFS = func(root *TreeNode, max_num int) {
		if root.Val >= max_num {
			count++
		}

		if root.Left != nil {
			DFS(root.Left, max(max_num, root.Val))
		}

		if root.Right != nil {
			DFS(root.Right, max(max_num, root.Val))
		}
	}
	DFS(root, root.Val)
	return count
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
