package running_sum

func runningSum(nums []int) []int {
	var sum int
	for i, v := range nums {
		sum += v
		nums[i] = sum
	}
	return nums
}
