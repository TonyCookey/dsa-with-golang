package lowest_common_ancestor_bst

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans, _ := LCA(root, p, q)
	return ans
}

func LCA(root, p, q *TreeNode) (*TreeNode, int) {
	if root == nil {
		return root, 0
	}
	leftNode, left := LCA(root.Left, p, q)

	rightNode, right := LCA(root.Right, p, q)

	fmt.Println(root, left, right)

	if left == 2 {
		return leftNode, 2
	}
	if right == 2 {
		return rightNode, 2
	}

	if left == 1 && right == 1 {
		return root, 2
	}

	if (root == p || root == q) && (left == 1 || right == 1) {
		return root, 2
	}

	if (root == p || root == q) || (left == 1 || right == 1) {
		return root, 1
	}
	if root == p || root == q {
		return root, 1
	}

	return root, 0
}
