package node_depth

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	var sum int

	var DFS func(root *BinaryTree, depth int)

	DFS = func(root *BinaryTree, depth int) {
		sum += depth
		if root.Left != nil {
			DFS(root.Left, depth+1)
		}

		if root.Right != nil {
			DFS(root.Right, depth+1)
		}
	}
	DFS(root, 0)
	return sum
}
