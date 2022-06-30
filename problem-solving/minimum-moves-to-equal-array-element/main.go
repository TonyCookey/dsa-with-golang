package minimum_moves_to_equal_array_element

import "sort"

func minMoves1(nums []int) int {
	sort.Ints(nums)
	var sum int

	for _, num := range nums {
		sum += num - nums[0]
	}

	return sum
}
