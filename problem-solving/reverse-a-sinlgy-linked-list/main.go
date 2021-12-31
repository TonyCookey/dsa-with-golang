package main

import (
	"fmt"
)

// Node struct - represents a node in a linked list
// see the data-structures/linked-list/singly-linked-list for the full set of methods
type Node struct {
	value int
	next  *Node
}

//LinkedList struct - contains a head and a length
type LinkedList struct {
	head   *Node
	length int
}

// append - add a node to the end of the linked list
func (l *LinkedList) append(value int) {
	node := Node{}
	node.value = value

	if l.length == 0 {
		l.head = &node
		l.length++
		//l.tail = &node
		return
	}
	// ptr - address of the starting/current point through the iteration
	ptr := l.head
	for i := 0; i < l.length; i++ {
		if ptr.next == nil {
			ptr.next = &node
			l.length++
			return
		}
		ptr = ptr.next
	}
}

// print - prints all the nodes in the linkedlist
func (l *LinkedList) print() {
	if l.length == 0 {
		fmt.Println("No Nodes in the Linked list")
	}
	slice := []int{}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		slice = append(slice, ptr.value)
		// using de-referencing and accessing a pointer value
		//slice = append(slice, (*ptr).value)
		ptr = ptr.next
	}
	fmt.Println("LinkedList", slice)
}

//reverseUsingIteration - reverse a singly linked list using iteration
func (l *LinkedList) reverseUsingIteration() {
	if l.length == 0 {
		return
	}
	var current, prev, next *Node

	//set the current to the head of the linked list
	current = l.head

	//traverse through the linked list until current is nil
	for current != nil {
		next = current.next
		current.next = prev

		//preparing values for the next iteration - i.e moving the cursors - current and prev
		prev = current
		current = next
	}
	l.head = prev
	l.print()
}
func main() {
	linkedlist := LinkedList{}

	linkedlist.append(4)
	linkedlist.append(5)
	linkedlist.append(18)
	linkedlist.append(12)
	linkedlist.append(8)
	linkedlist.append(2)
	linkedlist.append(3)
	linkedlist.append(1)
	linkedlist.print()
	linkedlist.reverseUsingIteration()
}
