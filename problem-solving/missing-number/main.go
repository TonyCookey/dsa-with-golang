package missing_number

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
func missingNumber2(nums []int) int {
	var sum int
	expectedSum := len(nums) * (len(nums) + 1) / 2
	for _, v := range nums {
		sum += v
	}
	return expectedSum - sum
}
