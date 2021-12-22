package main

import "fmt"

func main() {
	fmt.Println(fibonacciIterative(7))
	fmt.Println(fibonacciIterativeArray(7))
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
func fibonacciIterativeArray(n int) int {
	xi := []int{0, 1}

	for i := 2; i <= n; i++ {
		xi = append(xi, xi[i-2]+xi[i-1])
	}
	return xi[n]
}

func fibonacciRecursive(n int) int {
	if n <= 2 {
		return 1
	}
	sum := fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
	return sum
}
