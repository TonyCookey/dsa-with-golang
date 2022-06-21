package find_all_numbers_disappeared_in_an_array

import "sort"

func findDisappearedNumbers(nums []int) []int {
	var result []int
	sort.Ints(nums)

	k := len(nums)
	i := len(nums) - 1
	for k > 0 {
		v := nums[i]
		if v != k {
			result = append([]int{k}, result...)
			k--
			continue
		}
		if i == 0 {
			k--
			continue
		}

		if v == k {
			if nums[i] != nums[i-1] {
				k--
			}
			i--
		}
	}
	return result
}
