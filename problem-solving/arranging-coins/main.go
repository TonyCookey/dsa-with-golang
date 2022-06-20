package arranging_coins

func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	var i int
	m := n
	for i = 1; i < m; i++ {
		if n-i < 0 {
			break
		}
		n -= i
	}
	return i - 1
}
