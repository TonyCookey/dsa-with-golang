package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	result := []int{0}
	low, high := prices[0], prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < low {
			if high-low > result[len(result)-1] {
				result = append(result, high-low)
			}
			low = prices[i]
			high = prices[i]
		}
		if prices[i] > high {
			high = prices[i]
			if high-low > result[len(result)-1] {
				result = append(result, high-low)
			}
		}
	}
	return result[len(result)-1]
}
