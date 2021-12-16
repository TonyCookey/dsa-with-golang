package main

import "fmt"

type Node struct {
	value string
	next  *Node
}
type Stack struct {
	top    *Node
	bottom *Node
	length int
}

func (s *Stack) peek() string {
	fmt.Println(s.top.value)
	return s.top.value
}
func (s *Stack) isEmpty() bool {
	fmt.Println(s.length == 0)
	return s.length == 0
}
func (s *Stack) pop() {
	fmt.Println(s.bottom.value)
	if s.length == 1 {
		s.bottom = nil
		fmt.Println(s.bottom)
	}
	s.top = s.top.next
	s.length--
	return
}
func (s *Stack) push(value string) {
	newNode := Node{
		value: value,
		next:  s.top,
	}
	if s.length == 0 {
		s.bottom = &newNode
	}
	s.top = &newNode
	s.length++
	return
}

func main() {
	stack := Stack{}
	stack.push("Tony")
	stack.push("Cookey")
	stack.push("Starboy")
	stack.peek()
	fmt.Println(stack)
	stack.pop()
	stack.pop()
	stack.pop()
	stack.isEmpty()
	fmt.Println(stack)
}
