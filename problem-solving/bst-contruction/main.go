package main

//Do not edit the class below except for
//the insert, contains, and remove methods.
//Feel free to add new properties and methods
//to the class.
type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	// Write your code here.

	head := tree

	for head != nil {
		if value < head.Value {
			if head.Left == nil {
				head.Left = &BST{
					Value: value,
				}
				break
			}
			head = head.Left
		} else {
			if head.Right == nil {
				head.Right = &BST{
					Value: value,
				}
				break
			}
			head = head.Right
		}
	}
	// Do not edit the return statement of this method.

	return tree
}

func (tree *BST) Contains(value int) bool {

	for tree != nil {
		if tree.Value == value {
			return true
		}
		if tree.Value >= value {
			tree = tree.Left
		} else {
			tree = tree.Right
		}
	}
	return false
}

func (tree *BST) Remove(value int) *BST {
	// Write your code here.
	// Do not edit the return statement of this method.
	if tree == nil || (tree.Left == nil && tree.Right == nil) {
		return nil
	}
	return tree
}
