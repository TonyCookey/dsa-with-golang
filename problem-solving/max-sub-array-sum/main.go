package main

import "fmt"

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//[5,4,-1,7,8]
	fmt.Println(maxSubArray(input))
}
func maxSubArray(nums []int) int {
	// sum  & maxSum
	// length of nums
	// contiguous : that follow each other
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
