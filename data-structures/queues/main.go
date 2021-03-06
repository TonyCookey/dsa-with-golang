package main

import "fmt"

type Node struct {
	value string
	next  *Node
}
type Queue struct {
	first  *Node
	last   *Node
	length int
}

// peek - checks the first value in the queue
func (q *Queue) peek() string {
	fmt.Println(q.first.value)
	return q.first.value
}

// isEmpty - checks if the queue is empty
func (q *Queue) isEmpty() bool {
	fmt.Println(q.length == 0)
	return q.length == 0
}

// dequeue - remove from the queue - FIFO
func (q *Queue) dequeue() {
	if q.length == 1 {
		q.last = nil
	}
	q.first = q.first.next
	q.length--
	return
}

// enqueue - add to the queue
func (q *Queue) enqueue(value string) {
	newNode := Node{
		value: value,
	}
	if q.length == 0 {
		q.first = &newNode
		q.last = &newNode
	}
	q.last.next = &newNode
	q.last = &newNode
	q.length++
	return
}

func main() {
	queue := Queue{}
	queue.enqueue("Tony")
	queue.enqueue("Cookey")
	queue.enqueue("Starboy")
	fmt.Println(queue)
	queue.dequeue()
	queue.dequeue()
	fmt.Println(queue)
	queue.dequeue()
	fmt.Println(queue)
}
