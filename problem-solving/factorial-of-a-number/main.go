package main

import "fmt"

func main() {
	fmt.Println(findFactorialUsingRecursion(5))
}

// 5
func findFactorialUsingRecursion(x int) int {
	if x < 1 {
		return 1
	}
	sum := x * findFactorialUsingRecursion(x-1)
	return sum
}
