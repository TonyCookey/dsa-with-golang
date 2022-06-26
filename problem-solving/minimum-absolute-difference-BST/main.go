package minimum_absolute_difference_BST

import (
	"golang.org/x/tools/container/intsets"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference2(root *TreeNode) int {
	var inOrderDFS func(root *TreeNode)
	prev := -1
	min := intsets.MaxInt

	inOrderDFS = func(root *TreeNode) {
		if root != nil {
			inOrderDFS(root.Left)

			if prev != -1 && root.Val-prev < min {
				min = root.Val - prev
			}
			prev = root.Val
			inOrderDFS(root.Right)
		}
	}
	inOrderDFS(root)
	return min
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
