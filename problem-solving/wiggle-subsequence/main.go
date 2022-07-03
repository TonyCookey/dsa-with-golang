package wiggle_subsequence

func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	var diff int
	count := 1

	for i := 1; i < len(nums); i++ {
		if diff > 0 && nums[i]-nums[i-1] < 0 {
			count += 1
			diff = nums[i] - nums[i-1]
		}
		if diff < 0 && nums[i]-nums[i-1] > 0 {
			count += 1
			diff = nums[i] - nums[i-1]
		}
		if diff == 0 && nums[i]-nums[i-1] != 0 {
			count += 1
			diff = nums[i] - nums[i-1]
		}
	}
	return count
}
