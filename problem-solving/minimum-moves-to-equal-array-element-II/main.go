package minimum_moves_to_equal_array_element_II

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	mid := nums[len(nums)/2]
	var minMoves int

	for _, v := range nums {
		minMoves += abs(mid - v)
	}
	return minMoves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
