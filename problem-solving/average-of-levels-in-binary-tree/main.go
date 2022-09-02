package average_of_levels_in_binary_tree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := []*TreeNode{root}
	var ans []float64
	for len(queue) != 0 {
		n, sum := len(queue), 0
		for i := 0; i < n; i++ {
			pop := queue[0]
			queue = queue[1:]
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
			sum += pop.Val
		}
		ans = append(ans, float64(sum)/float64(n))
	}
	return ans
}
