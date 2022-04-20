package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var first, second, prev *TreeNode

func recoverTree(root *TreeNode) {
	first, second, prev = nil, nil, nil
	inOrderTraversal(root)
	fmt.Println(first, second, prev)

	temp := first.Val
	first.Val = second.Val
	second.Val = temp
}
func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left)

	if prev != nil && prev.Val > root.Val {
		if first == nil {
			first = prev
		}
		second = root
	}

	prev = root

	inOrderTraversal(root.Right)
}
