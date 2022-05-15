package depeest_leaves_sum

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	var sum int
	for len(queue) != 0 {
		queueLen := len(queue)
		sum = 0
		for i := 0; i < queueLen; i++ {
			curr := queue[0]
			queue = queue[1:]
			sum += curr.Val
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}
	return sum
}
