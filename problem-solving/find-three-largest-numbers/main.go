package main

import (
	"golang.org/x/tools/container/intsets"
)

func FindThreeLargestNumbers(array []int) []int {
	// Write your code here.

	first := intsets.MinInt
	second := intsets.MinInt
	third := intsets.MinInt

	for i := 0; i < len(array); i++ {
		if array[i] > first {
			third = max(third, second)
			second = max(second, first)
			first = array[i]
			continue
		} else if array[i] > second {
			third = max(third, second)
			second = array[i]
			continue
		} else if array[i] > third {
			third = array[i]
		}
	}
	return []int{third, second, first}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
