package min_cost_climbing_stairs

func minCostClimbingStairs(cost []int) int {
	dp := make(map[int]int)
	return min(recursive(cost, 0, dp), recursive(cost, 1, dp))
}
func recursive(cost []int, index int, dp map[int]int) int {
	if index > len(cost)-1 {
		return 0
	}
	if index+1 > len(cost)-1 {
		return cost[index]
	}
	if v, ok := dp[index]; ok {
		return v
	}

	dp[index] = min(cost[index]+recursive(cost, index+1, dp), cost[index]+recursive(cost, index+2, dp))
	return dp[index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
