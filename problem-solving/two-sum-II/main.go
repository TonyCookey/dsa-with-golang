package two_sum_II

func twoSum(numbers []int, target int) []int {
	seen := make(map[int]int)

	for i, v := range numbers {
		if seen[target-v] > 0 {
			return []int{seen[target-v], i + 1}
		} else {
			seen[v] = i + 1
		}
	}
	return nil
}
