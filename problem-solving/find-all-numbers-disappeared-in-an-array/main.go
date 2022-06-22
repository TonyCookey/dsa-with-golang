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
func findDisappearedNumbersOptimal(nums []int) []int {
	var result []int

	for _, v := range nums {
		nums[abs(v)-1] = -abs(nums[abs(v)-1])
	}

	for i, v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
