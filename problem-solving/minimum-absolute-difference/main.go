package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(findMinimumAbsoluteDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}
func findMinimumAbsoluteDifference(arr []int) [][]int {
	// sort the arr
	sort.Ints(arr)

	// keep track of the minimum variable
	difference := make(map[int][][]int)
	arrLength := len(arr)
	minValue := arr[arrLength-1]

	for i := 0; i < arrLength-1; i++ {
		current := arr[i]
		next := arr[i+1]
		if value, ok := difference[next-current]; ok {
			difference[next-current] = append(value, []int{current, next})
		} else {
			difference[next-current] = append([][]int{}, []int{current, next})
		}
		minValue = int(math.Min(float64(minValue), float64(next-current)))
	}
	return difference[minValue]

}
