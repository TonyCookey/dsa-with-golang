package remove_duplicates_from_sorted_array

import (
	"golang.org/x/tools/container/intsets"
)

func removeDuplicates(nums []int) int {
	exists := make(map[int]bool)
	var minValuePointer int
	for i := 0; i < len(nums); i++ {
		if nums[i] == intsets.MinInt {
			continue
		}

		if exists[nums[i]] {
			if minValuePointer == 0 {
				minValuePointer = i
			}
			nums[i] = intsets.MinInt
		} else {
			exists[nums[i]] = true
			if minValuePointer != 0 {
				nums[minValuePointer] = nums[i]
				nums[i] = intsets.MinInt
				minValuePointer++
			}
		}
	}
	return len(exists)
}
