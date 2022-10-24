package branch_sums

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	var result []int
	var DFS func(root *BinaryTree, sum int)

	DFS = func(root *BinaryTree, sum int) {
		if root == nil {
			return
		}
		sum += root.Value
		if root.Left == nil && root.Right == nil {
			result = append(result, sum)
		}

		DFS(root.Left, sum)
		DFS(root.Right, sum)
	}
	DFS(root, 0)
	return result
}
