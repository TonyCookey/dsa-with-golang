package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	var maxProfit int
	left, right := 0, 1

	for right < len(prices) {
		if prices[left] < prices[right] {
			maxProfit = max(maxProfit, prices[right]-prices[left])
		} else {
			left = right
		}
		right++

	}
	return maxProfit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
