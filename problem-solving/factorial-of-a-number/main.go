package main

import "fmt"

func main() {
	fmt.Println(findFactorialUsingRecursion(5))
	fmt.Println(findFactorialUsingIteration(5))
}

// 5
func findFactorialUsingRecursion(x int) int {
	if x < 1 {
		return 1
	}
	return x * findFactorialUsingRecursion(x-1)

}
func findFactorialUsingIteration(x int) int {
	var sum = 1
	for x > 0 {
		sum *= x
		x--
	}
	return sum
}
