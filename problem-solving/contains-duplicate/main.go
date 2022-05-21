package contains_duplicate

import "sort"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	// sort to make faster, duplicates will be together
	sort.Ints(nums)
	for _, val := range nums {
		if seen[val] {
			return true
		}
		seen[val] = true
	}
	return false
}
