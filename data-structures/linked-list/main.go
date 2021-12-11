package main

//Node struct - represents a node in a linked list
type Node struct {
	value int
	next  *Node
}

//LinkedList struct - contains a head and a length
type LinkedList struct {
	head   *Node
	length int
	tail   *Node
}
