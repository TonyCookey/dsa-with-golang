package leaf_similar_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := getLeaves(root1, []int{})
	leaves2 := getLeaves(root2, []int{})

	if len(leaves1) != len(leaves2) {
		return false
	}
	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}
func getLeaves(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	if root.Left == nil && root.Right == nil {
		arr = append(arr, root.Val)
		return arr
	}
	arr = getLeaves(root.Left, arr)
	arr = getLeaves(root.Right, arr)
	return arr
}
