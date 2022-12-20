package longest_common_subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)

	for index, _ := range dp {
		dp[index] = make([]int, len(text2)+1)
	}
	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[0][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
