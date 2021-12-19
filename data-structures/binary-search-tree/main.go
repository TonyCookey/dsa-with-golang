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
		if currentNode.value > value {
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

// lookup- finds a node using its value
func (b *BinarySearchTree) lookup(value int) *Node {
	currentNode := b.root
	for currentNode != nil {
		if currentNode.value == value {
			return currentNode
		}
		if currentNode.value < value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}
	return nil
}

func main() {
	bst := BinarySearchTree{}
	bst.insert(10)
	bst.insert(8)
	bst.insert(6)
	bst.insert(14)
	bst.insert(2)
	bst.insert(9)
	bst.insert(7)
	bst.insert(12)
	bst.insert(15)
	bst.insert(11)
	fmt.Println(bst.lookup(2))
	fmt.Println(bst)
}
