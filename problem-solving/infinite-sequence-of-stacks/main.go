package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

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
	if A[len(A)-1] > 1 {
		A[len(A)-1] = 1
	}

	return sum(A)

}
func sum(intSlice []int) int {
	sum := 0
	for i := 0; i < len(intSlice); i++ {
		sum += intSlice[i]
	}
	return sum
}
