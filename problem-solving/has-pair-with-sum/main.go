package main

import "fmt"

func main() {
	sliceOfInts := []int{7, 4, 5, 2, 3, 4, 1, 8}
	sum := 10
	fmt.Println(usingBruteForce(sliceOfInts, sum))
	fmt.Println(hasPairWithSum(sliceOfInts, sum))
}

// Solving hasPairWithSum using maps - a more efficient approach
func hasPairWithSum(slice []int, sum int) bool {
	//Time Complexity : O(n)
	// used to store the complements - (sum - slice[i]), which is what is needed for a pair sum
	complements := map[int]bool{}
	for i := 0; i < len(slice); i++ {
		fmt.Println("complements", complements)
		if complements[slice[i]] {
			return true
		}
		complements[sum-slice[i]] = true
	}
	return false
}

// Solving hasPairWithSum using the brute force approach
func usingBruteForce(slice []int, sum int) bool {
	//Time Complexity : O(n^2)
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i]+slice[j] == sum {
				return true
			}
		}
	}
	return false
}
