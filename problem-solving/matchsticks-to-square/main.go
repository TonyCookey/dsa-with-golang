package matchsticks_to_square

import (
	"fmt"
	"sort"
)

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}
	sumMatchsticks := sum(matchsticks)
	length := sumMatchsticks / 4
	sides := make([]int, 4)
	if float64(sumMatchsticks)/4 != float64(length) {
		return false
	}
	sort.SliceStable(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	var backtrack func(i int) bool

	backtrack = func(i int) bool {
		if i == len(matchsticks) {
			return true
		}
		for j := 0; j < 4; j++ {
			if sides[j]+matchsticks[i] <= length {
				sides[j] += matchsticks[i]
				if backtrack(i + 1) {
					return true
				}
				sides[j] -= matchsticks[i]
			}
		}
		return false
	}
	fmt.Println(matchsticks, length)
	return backtrack(0)
}

func sum(x []int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
