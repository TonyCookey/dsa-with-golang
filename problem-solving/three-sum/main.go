package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := v + nums[left] + nums[right]

			if sum > 0 {
				right -= 1
			} else if sum < 0 {
				left += 1
			} else {
				result = append(result, []int{v, nums[left], nums[right]})
				left += 1
				for nums[left] == nums[left-1] && left < right {
					left += 1
				}
			}
		}
	}
	return result
}
