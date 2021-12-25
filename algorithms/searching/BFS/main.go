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

func (b *BinarySearchTree) breadthFirstSearch() []int {
	currentNode := b.root
	queue := []*Node{}
	list := []int{}
	queue = append(queue, currentNode)

	for len(queue) != 0 {
		currentNode = queue[0]
		list = append(list, currentNode.value)

		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
		queue = queue[1:]
	}
	return list
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
	//bst.insert(12)
	//bst.insert(15)
	//bst.insert(11)
	fmt.Println(bst.breadthFirstSearch())
}
