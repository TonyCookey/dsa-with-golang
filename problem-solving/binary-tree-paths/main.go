package binary_tree_paths

import "strconv"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}

	bptHelper(root, &result, "")

	return result
}

func bptHelper(root *TreeNode, result *[]string, str string) {
	if root == nil {
		return
	}

	if str != "" {
		str += "->"
	}

	str += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*result = append(*result, str)
		return
	}

	bptHelper(root.Left, result, str)

	bptHelper(root.Right, result, str)

	return
}
