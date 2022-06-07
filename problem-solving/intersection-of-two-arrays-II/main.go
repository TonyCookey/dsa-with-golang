package intersection_of_two_arrays_II

func intersect(nums1 []int, nums2 []int) []int {
	seen := make(map[int]int)
	var result []int

	for _, v := range nums1 {
		seen[v]++
	}

	for _, v := range nums2 {
		if seen[v] > 0 {
			result = append(result, v)
		}
		seen[v]--
	}
	return result
}
