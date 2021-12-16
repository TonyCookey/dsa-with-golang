package main

import "fmt"

type Stack struct {
	data []string
}

func (s *Stack) peek() string {
	return s.data[len(s.data)-1]
}
func (s *Stack) isEmpty() bool {
	fmt.Println(len(s.data) == 0)
	return len(s.data) == 0
}
func (s *Stack) pop() {
	if len(s.data) == 0 {
		return
	}
	//s.data = append([]string{}, s.data[:len(s.data) - 1]...)
	s.data = s.data[:len(s.data)-1]
	fmt.Println(s.data)
	return
}
func (s *Stack) push(value string) {
	s.data = append(s.data, value)
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
