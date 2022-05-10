package single_number

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)
	var curr int
	for i := 0; i < len(nums); {
		curr = nums[i]
		if i+1 >= len(nums) {
			break
		}
		if curr == nums[i+1] {
			i += 2
		} else {
			return curr
		}
	}
	return curr
}
