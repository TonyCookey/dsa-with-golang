package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func main() {
	fmt.Println(Solution([]int{4, 5, 0, 2, 1}))
}
func Solution(A []int) int {
	// write your code in Go 1.4
	A = append(A, 0)
	for i := 0; i < len(A)-1; i++ {

		if A[i] < 2 {
			continue
		}
		var x int = A[i] / 2

		if A[i]%2 == 0 {
			A[i] = 0
		} else {
			A[i] = 1
		}
		A[i+1] += x

	}
	ans := sumSlice(A)

	back := A[len(A)-1]

	for back > 0 {
		if back%2 == 0 {
			back /= 2
		} else {
			ans++
			back--
			back /= 2
		}
	}

	return ans

}
func sumSlice(intSlice []int) int {
	sum := 0
	for i := 0; i < len(intSlice)-1; i++ {
		sum += intSlice[i]
	}
	return sum
}
