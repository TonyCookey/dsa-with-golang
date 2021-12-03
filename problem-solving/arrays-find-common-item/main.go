package main

import (
	"fmt"
	"time"
)

func main() {
	arr1 := []string{"a", "b", "z", "e", "r", "o", "p", "r", "w", "k"}
	arr2 := []string{"l", "h", "x", "t", "a", "q", "s"}
	fmt.Println("Using brute force approach:  ", usingBruteForce(arr1, arr2))
	fmt.Println("Using hash tables: ", usingHashTable(arr1, arr2))
}
func usingBruteForce(arr1, arr2 []string) bool {
	start := time.Now()
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				duration := time.Since(start)
				fmt.Println("Brute force Time", duration.Nanoseconds())
				return true
			}
		}
	}

	return false
}
func usingHashTable(arr1, arr2 []string) bool {
	// convert the first array to a hash map/table
	// use the hash map/table to check against the second array to see if that field exists
	// thereby reducing the time complexity from O(n^2) to O(a+b)
	start := time.Now()

	mapStore := map[string]bool{}

	for i := 0; i < len(arr1); i++ {
		if mapStore[arr1[i]] == false {
			mapStore[arr1[i]] = true
		}
	}
	for j := 0; j < len(arr2); j++ {
		if mapStore[arr2[j]] {
			duration := time.Since(start)
			fmt.Println("hash tables time", duration.Nanoseconds())
			return true
		}
	}
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())
	return false
}
