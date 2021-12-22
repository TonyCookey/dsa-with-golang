package main

import "fmt"

func main() {
	fmt.Println(fibonacciIterative(12))
}

func fibonacciIterative(n int) int {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
	current := 1
	prev := 0
	next := 0
	index := 1

	for n > index {
		next = current + prev
		prev = current
		current = next
		index++
	}
	return current
}
