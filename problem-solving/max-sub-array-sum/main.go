package main

import "fmt"

func main() {
	input := []int{-2}
	fmt.Println(maxSubArray(input))
}
func maxSubArray(nums []int) int {
	// sum  & maxsum
	// length of nums
	// contiguous : that follow each other

	numsLength := len(nums)
	if numsLength == 0 || nums == nil {
		return 0
	}

	var maxSum int = nums[0]
	var sum int = nums[0]

	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if maxSum < sum {
			maxSum = sum
		}
	}
	return maxSum

}
