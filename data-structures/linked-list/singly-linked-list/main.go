package main

import (
	"errors"
	"fmt"
	"log"
)

//Node struct - represents a node in a linked list
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

func (l *LinkedList) getLastNode() *Node {
	return l.getAtPosition(l.length - 1)
}

// getAtPosition - returns the node at the specified index
func (l *LinkedList) getAtPosition(index int) *Node {
	ptr := l.head
	if index == 0 {
		return ptr
	}
	if index < 0 {
		return nil
	}
	if index > (l.length - 1) {
		return nil
	}
	for i := 0; i < index; i++ {
		ptr = ptr.next
	}
	return ptr
}

// print - prints all the nodes in the linkedlist
func (l *LinkedList) print() {
	if l.length == 0 {
		fmt.Println("No Nodes in the Linked list")
	}
	slice := []Node{}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		slice = append(slice, *ptr)
		ptr = ptr.next
	}
	fmt.Println("LinkedList", slice)
}

// insertAtPosition : inserts a node at a specified index
func (l *LinkedList) insertAtPosition(index int, value int) {

	if index == 0 {
		l.head = &Node{
			value: value,
			next:  l.head,
		}
		l.length++
		return
	}
	if index < 0 {
		return
	}
	if index > (l.length - 1) {
		l.append(value)
	}
	leaderNode := l.getAtPosition(index - 1)
	currentNode := leaderNode.next
	leaderNode.next = &Node{
		value: value,
		next:  currentNode,
	}
	l.length++
}

// prepend - inserts at the beginning of the linked list
func (l *LinkedList) prepend(value int) {
	l.insertAtPosition(0, value)
}
func (l *LinkedList) searchByValue(value int) *Node {
	if l.length == 0 {
		return nil
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		if ptr.value == value {
			return ptr
		}
		ptr = ptr.next
	}
	return nil
}
func (l *LinkedList) deleteAtPosition(index int) error {
	if l.length == 0 {
		return errors.New("no nodes in the linkedList")
	}
	if index == 0 {
		l.head = l.head.next
		l.length--
	}
	if index < 0 {
		return errors.New("invalid index: no negative numbers")
	}
	if index > (l.length - 1) {
		return errors.New("invalid index: index greater than length of linked list")
	}
	leaderNode := l.getAtPosition(index - 1)
	currentNode := leaderNode.next
	leaderNode.next = currentNode.next
	l.length--

	return nil
}
func (l *LinkedList) deleteNodeUsingValue(value int) error {
	if l.length == 0 {
		return errors.New("no nodes in the linkedList")
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		if ptr.value == value {
			if i == 0 {
				l.head = ptr.next
				l.length--
				return nil
			}
			leaderNode := l.getAtPosition(i - 1)
			leaderNode.next = ptr.next
			l.length--
			return nil
		}
		ptr = ptr.next
	}
	return nil

}
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
	linkedlist.prepend(50)
	linkedlist.getAtPosition(1)
	fmt.Println(linkedlist.searchByValue(5))
	linkedlist.print()
	linkedlist.insertAtPosition(2, 10)
	linkedlist.print()
	err := linkedlist.deleteAtPosition(1)
	if err != nil {
		log.Fatalln(err)
	}
	linkedlist.print()
	err = linkedlist.deleteNodeUsingValue(50)
	if err != nil {
		log.Fatalln(err)
	}
	linkedlist.print()
}
