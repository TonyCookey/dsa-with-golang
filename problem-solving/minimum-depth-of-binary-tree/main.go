package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	var depth int

	for len(queue) > 0 {
		queue_length := len(queue)
		depth++

		for i := 0; i < queue_length; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr.Right == nil && curr.Left == nil {
				return depth
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}

		}
	}
	return depth
}

func minDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := minDepthDFS(root.Left)
	rightDepth := minDepthDFS(root.Right)

	return min(leftDepth, rightDepth) + 1
}

func min(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}

	if x < y {
		return x
	}
	return y
}
