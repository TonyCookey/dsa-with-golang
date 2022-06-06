package intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	var result []int

	seen := make(map[int]bool)
	ans := make(map[int]bool)

	for _, val := range nums1 {
		seen[val] = true
	}

	for _, val := range nums2 {
		if seen[val] && !ans[val] {
			ans[val] = true
			result = append(result, val)
		}
	}
	return result
}
