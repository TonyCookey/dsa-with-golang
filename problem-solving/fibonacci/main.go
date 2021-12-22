package main

import "fmt"

func main() {
	fmt.Println(fibonacciIterative(7))
	fmt.Println(fibonacciRecursive(7))
}

func fibonacciIterative(n int) int {
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
func fibonacciRecursive(n int) int {
	if n <= 2 {
		return 1
	}
	sum := fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
	fmt.Println("sum:", n-1, "-", n-2, "=", sum)
	return sum
}
