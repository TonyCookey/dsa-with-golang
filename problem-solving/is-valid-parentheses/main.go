package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	stack := []uint8{s[0]}

	for i := 1; i < len(s); i++ {

		v := s[i]

		if v == ')' || v == '}' || v == ']' {
			if len(stack) == 0 {
				return false
			}
			top_stack := stack[len(stack)-1]
			if v == ')' && top_stack != '(' || v == '}' && top_stack != '{' || v == ']' && top_stack != '[' {
				fmt.Println("3", top_stack, v)
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}
	fmt.Println("stack", stack)

	return len(stack) == 0
}
