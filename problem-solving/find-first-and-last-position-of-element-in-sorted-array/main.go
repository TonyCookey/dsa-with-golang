package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			for nums[left] != target || nums[right] != target {
				if nums[left] != target {
					left++
				}
				if nums[right] != target {
					right--
				}
			}
			return []int{left, right}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}
