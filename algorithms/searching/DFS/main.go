package main

import "fmt"

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

func (b *BinarySearchTree) DFSPostOrder(node *Node, list []int) []int {
	if node.left != nil {
		list = b.DFSPostOrder(node.left, list)
	}
	if node.right != nil {
		list = b.DFSPostOrder(node.right, list)
	}
	list = append(list, node.value)
	return list
}

func main() {
	bst := BinarySearchTree{}
	bst.insert(9)
	bst.insert(4)
	bst.insert(20)
	bst.insert(1)
	bst.insert(6)
	bst.insert(15)
	bst.insert(170)
	fmt.Println(bst.DFSPostOrder(bst.root, []int{}))
}
