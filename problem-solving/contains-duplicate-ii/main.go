package contains_duplicate_ii

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)

	for i, v := range nums {
		if seen[v] > 0 && abs(seen[v]-(i+1)) <= k {
			return true
		} else {
			seen[v] = i + 1
		}
	}
	return false
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
