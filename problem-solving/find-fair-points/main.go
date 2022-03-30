package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
	// write your code in Go 1.4
	count := 0
	sumA := sumSlice(A)
	sumB := sumSlice(B)

	sumA1 := 0
	sumB1 := 0

	for k := 0; k < len(A)-1; k++ {
		sumA1 += A[k]
		sumB1 += B[k]
		if sumA1 == sumA-sumA1 && sumB1 == sumB-sumB1 && sumA1 == sumB1 {
			count++
		}
	}

	return count
}
func sumSlice(intSlice []int) int {
	sum := 0
	for i := 0; i < len(intSlice); i++ {
		sum += intSlice[i]
	}
	return sum
}
