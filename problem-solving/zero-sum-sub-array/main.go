package main

func ZeroSumSubarray(nums []int) bool {
	sumSet := make(map[int]bool)
	currSum := 0

	for _, num := range nums {
		currSum += num
		if currSum == 0 || sumSet[currSum] {
			return true
		}
		sumSet[currSum] = true
	}
	return false
}

func ZeroSumSubarrayBruteForce(nums []int) bool {
	// Write your code here.
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == 0 {
				return true
			}
		}
	}
	return false
}
