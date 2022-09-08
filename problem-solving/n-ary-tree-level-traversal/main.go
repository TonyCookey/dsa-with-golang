package n_ary_tree_level_traversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	var currentNode *Node
	queue := []*Node{root}

	var list [][]int

	for len(queue) != 0 {
		length := len(queue)
		var sub_list []int
		for i := 0; i < length; i++ {
			currentNode = queue[0]

			sub_list = append(sub_list, currentNode.Val)

			for _, node := range currentNode.Children {
				if node != nil {
					queue = append(queue, node)
				}
			}
			queue = queue[1:]
		}
		list = append(list, sub_list)
	}
	return list
}
