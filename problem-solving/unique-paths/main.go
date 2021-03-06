package unique_paths

// Dynamic Programming
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := m - 1; i > -1; i-- {
		dp[i] = make([]int, n)
		for j := n - 1; j > -1; j-- {
			if i == m-1 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}
