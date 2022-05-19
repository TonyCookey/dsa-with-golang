package squares_of_sorted_array

import (
	"math"
	"sort"
)

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = int(math.Pow(float64(nums[i]), 2))
	}
	sort.Ints(nums)
	return nums
}
