package minimum_moves_to_equal_array_element

import (
	"golang.org/x/tools/container/intsets"
	"sort"
)

func minMoves1(nums []int) int {
	sort.Ints(nums)
	var sum int

	for _, num := range nums {
		sum += num - nums[0]
	}

	return sum
}
func minMoves2(nums []int) int {
	var sum int
	minVal := intsets.MaxInt

	for _, num := range nums {
		sum += num
		if num < minVal {
			minVal = num
		}
	}

	return sum - len(nums)*minVal
}
