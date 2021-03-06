package search_index_position

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var mid int
	for low <= high {
		// calculate mid point
		mid = low + ((high - low) / 2)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
