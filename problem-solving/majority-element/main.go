package majority_element

func majorityElement(nums []int) int {
	seen := make(map[int]int)

	for _, v := range nums {
		if seen[v] > 0 {
			seen[v]++
		} else {
			seen[v] = 1
		}
		if seen[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}
