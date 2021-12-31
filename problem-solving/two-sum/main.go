package main

import "fmt"

func main() {
	sliceOfInts := []int{7, 9, 5, 2, 1}
	sum := 8
	fmt.Println(twoSum(sliceOfInts, sum))
	fmt.Println(twoSumCleaner(sliceOfInts, sum))
	fmt.Println(usingBruteForce(sliceOfInts, sum))
}

// twoSum using maps - a more efficient approach
func twoSum(nums []int, target int) []int {
	// Time Complexity: O(n)
	// get the indices of the pair that sum up to target
	complements := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if complements[nums[i]] != 0 {
			return []int{complements[nums[i]] - 1, i}
		}
		complements[target-nums[i]] = i + 1
	}
	return nil
}

// twoSumCleaner - solving twoSum using maps/hash maps - a more efficient and cleaner approach
func twoSumCleaner(nums []int, target int) []int {
	// Time Complexity: O(n)
	// get the indices of the pair that sum up to target
	complements := map[int]int{}
	for index, value := range nums {
		if _, ok := complements[value]; ok {
			return []int{complements[value], index}
		}
		complements[target-value] = index
	}
	return nil
}

// usingBruteForce - solving using the brute force approach
func usingBruteForce(nums []int, target int) []int {
	//Time Complexity : O(n^2)
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
