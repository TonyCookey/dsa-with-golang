package binary_tree_paths

import "strconv"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func binaryTreePaths(root *TreeNode) []string {

	if root == nil {
		return nil
	}
	var result []string

	bptHelper(root, &result, "")

	return result
}

func bptHelper(root *TreeNode, result *[]string, str string) {

	if str != "" {
		str += "->"
	}
	str += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*result = append(*result, str)
		return
	}

	if root.Left != nil {
		bptHelper(root.Left, result, str)
	}

	if root.Right != nil {
		bptHelper(root.Right, result, str)
	}
	return
}
