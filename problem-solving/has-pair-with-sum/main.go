package main

import "fmt"

func main() {
	sliceOfInts := []int{2, 4, 5, 7, 8}
	sum := 10
	fmt.Println(hasPairWithSum(sliceOfInts, sum))
}

// Solving hasPairWithSum using the brute force approach
func hasPairWithSum(slice []int, sum int) bool {

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i]+slice[j] == sum {
				return true
			}
		}
	}
	return false
}
