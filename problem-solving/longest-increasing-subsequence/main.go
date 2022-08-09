package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	ans, n := 0, len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
