package main

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isUnivalTreeHelper(root, root.Val)

}

func isUnivalTreeHelper(root *TreeNode, key int) bool {

	if root == nil {
		return true
	}

	if root.Val != key {
		return false
	}
	return isUnivalTreeHelper(root.Left, key) && isUnivalTreeHelper(root.Right, key)

}
