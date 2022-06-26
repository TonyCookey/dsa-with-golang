package minimum_absolute_difference_BST

import (
	"golang.org/x/tools/container/intsets"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var sorted []int
	inOrderDFS(root, &sorted)

	min := intsets.MaxInt

	for i := 1; i < len(sorted); i++ {
		if min > (sorted[i] - sorted[i-1]) {
			min = sorted[i] - sorted[i-1]
		}
	}
	return min
}

func inOrderDFS(root *TreeNode, sorted *[]int) {
	if root == nil {
		return
	}

	inOrderDFS(root.Left, sorted)

	*sorted = append(*sorted, root.Val)

	inOrderDFS(root.Right, sorted)
}
