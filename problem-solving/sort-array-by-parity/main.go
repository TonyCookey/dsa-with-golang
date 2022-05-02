package sort_array_by_parity

func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1

	for i < j {
		if nums[i]%2 == 0 {
			i++
			continue
		} else {
			if nums[j]%2 == 0 {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				j--
				continue
			}
		}
	}
	return nums
}
