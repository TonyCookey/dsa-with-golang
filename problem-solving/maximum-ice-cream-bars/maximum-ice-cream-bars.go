package maximum_ice_cream_bars

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	var currCost, count int

	for _, cost := range costs {
		currCost += cost
		if currCost > coins {
			return count
		}
		count++
	}
	return count
}
