package maximum_ice_cream_bars

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	var curr_cost, count int

	for _, cost := range costs {
		curr_cost += cost
		if curr_cost > coins {
			return count
		}
		count++
	}
	return count
}
