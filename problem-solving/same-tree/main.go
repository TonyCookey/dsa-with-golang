package same_tree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if q == nil && p != nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) == true && isSameTree(p.Right, q.Right) == true
}
