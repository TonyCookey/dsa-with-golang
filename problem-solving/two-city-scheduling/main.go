package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoCitySchedulingCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}))

}

func twoCitySchedulingCost(costs [][]int) int {
	fmt.Println(costs)

	sort.Slice(costs, func(a, b int) bool {
		return costs[a][0]-costs[a][1]-(costs[b][0]-costs[b][1]) > 0 // using aCost-bCost to compare
	})
	fmt.Println(costs)

	res := 0
	for i := 0; i < len(costs)/2; i++ {
		res += costs[i][1] + costs[len(costs)-1-i][0] // adding bCost of first n and aCost of last n
	}
	return res
}
