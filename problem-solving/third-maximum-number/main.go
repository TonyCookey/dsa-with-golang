package third_maximum_number

import "sort"

func thirdMax(nums []int) int {
	sort.Ints(nums)
	max := nums[len(nums)-1]

	var i, j int
	for i, j = len(nums)-1, 0; i >= 0 && j < 2; i-- {
		if nums[i] < max {
			j++
			max = nums[i]
		}
	}
	if j < 2 {
		return nums[len(nums)-1]
	}

	return max
}
