package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		levelSize := len(queue)
		var level []int

		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			level = append(level, curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			queue = queue[1:]
		}
		result = append(result, level)
	}
	return result
}
