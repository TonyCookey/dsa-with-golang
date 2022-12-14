package house_robber

func rob(nums []int) int {
	var max_num int
	for i := len(nums) - 3; i >= 0; i-- {
		max_num = nums[i+2]
		for j := i + 2; j <= i+3 && j <= len(nums)-1; j++ {
			max_num = max(max_num, nums[j])
		}
		nums[i] += max_num
	}
	for _, val := range nums[:2] {
		max_num = max(max_num, val)
	}
	return max_num
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
