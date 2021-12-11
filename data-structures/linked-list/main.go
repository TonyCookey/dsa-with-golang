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

// append - add a node to the end of the linked list
func (l LinkedList) append(value int) {
	node := Node{}
	node.value = value

	if l.length == 0 {
		l.head = &node
		l.length++
		l.tail = &node
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
