package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans := root
	var traverse func(root *TreeNode) *TreeNode

	traverse = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		left := traverse(root.Left)
		right := traverse(root.Right)
		var found int
		var foundNode *TreeNode
		if left == p || left == q {
			found++
			foundNode = left
		}

		if right == p || right == q {
			found++
			foundNode = right
		}

		if root == p || root == q {
			found++
			foundNode = root
		}

		if found >= 2 {
			ans = root
		}
		return foundNode
	}

	traverse(root)
	return ans
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var traverse func(root *TreeNode) *TreeNode

	traverse = func(root *TreeNode) *TreeNode {
		if root == nil || root == p || root == q {
			return root
		}

		left := traverse(root.Left)
		right := traverse(root.Right)

		if left != nil && right != nil {
			return root
		}
		if left != nil {
			return left
		}
		return right
	}

	return traverse(root)
}
