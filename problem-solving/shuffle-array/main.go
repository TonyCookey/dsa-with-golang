package shuffle_array

func shuffle(nums []int, n int) []int {
	var result []int
	i, j := 0, n

	for i < n || j < len(nums) {
		if i < n {
			result = append(result, nums[i])
		}
		if j < len(nums) {
			result = append(result, nums[j])
		}
		i++
		j++
	}
	return result
}
