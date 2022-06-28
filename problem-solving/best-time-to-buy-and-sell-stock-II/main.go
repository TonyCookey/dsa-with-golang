package best_time_to_buy_and_sell_stock_II

func maxProfit(prices []int) int {
	var profit int

	low, high := 0, 1

	for high < len(prices) {
		if prices[low] < prices[high] {
			profit += prices[high] - prices[low]
		}
		low = high
		high++
	}
	return profit

}
