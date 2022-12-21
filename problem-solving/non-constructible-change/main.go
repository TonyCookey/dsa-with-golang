package non_constructible_change

import "sort"

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)

	changeCreated := 0

	for _, val := range coins {
		if val > changeCreated+1 {
			return changeCreated + 1
		}
		changeCreated += val
	}

	return changeCreated + 1
}
