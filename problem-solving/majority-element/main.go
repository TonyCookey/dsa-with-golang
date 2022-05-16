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
func majorityElement2(nums []int) int {
	var count, major int

	for _, val := range nums {
		if count == 0 {
			major = val
		}
		if major == val {
			count += 1
			continue
		}
		count -= 1
	}
	return major
}
