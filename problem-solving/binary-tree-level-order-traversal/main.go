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
	for len(queue) > 0 {
		subQueueLen := len(queue)

		var subResult []int
		for i := 0; i < subQueueLen; i++ {
			curr := queue[0]
			queue = queue[1:]
			subResult = append(subResult, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		result = append(result, subResult)
	}
	return result
}
