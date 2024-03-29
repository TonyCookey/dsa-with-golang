package main

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}
type BinarySearchTree struct {
	root *Node
}

// inserts values into the binary search tree
func (b *BinarySearchTree) insert(value int) {
	newNode := Node{
		value: value,
	}
	if b.root == nil {
		b.root = &newNode
		return
	}
	//set the current/pointer node
	currentNode := b.root
	for currentNode != nil {
		if currentNode.value < value {
			if currentNode.right == nil {
				currentNode.right = &newNode
				return
			} else {
				currentNode = currentNode.right
			}

		} else {
			if currentNode.left == nil {
				currentNode.left = &newNode
				return
			} else {
				currentNode = currentNode.left
			}
		}
	}
	return
}
func (b *BinarySearchTree) isValidBST(node *Node) bool {

	if node.left == nil && node.right == nil {
		return true
	}
	answer := BSTTraverseInorder(node, []int{})
	for i := 0; i < len(answer)-1; i++ {
		if answer[i] >= answer[i+1] {
			return false
		}
	}
	return true
}
func BSTTraverseInorder(node *Node, list []int) []int {

	if node.left != nil {
		list = BSTTraverseInorder(node.left, list)
	}

	list = append(list, node.value)
	if node.right != nil {
		list = BSTTraverseInorder(node.right, list)
	}
	return list
}

// isValidBSTOptimum - optimum way to validate a binary search tree
func isValidBSTOptimum(node *Node) bool {
	return validateHelper(node, math.MinInt64, math.MaxInt64)
}

//validateHelper - helper function to validate a BST
func validateHelper(node *Node, min, max int) bool {
	if node == nil {
		return true
	}
	if !(node.value >= min && node.value < max) {
		return false
	}
	return validateHelper(node.left, min, node.value) && validateHelper(node.right, node.value, max)
}

func main() {
	bst := BinarySearchTree{}
	bst.insert(9)
	bst.insert(4)
	bst.insert(20)
	bst.insert(1)
	bst.insert(8)
	bst.insert(15)
	bst.insert(70)
	fmt.Println(isValidBSTOptimum(bst.root))
}
