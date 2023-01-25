package main

import "math"

func FindThreeLargestNumbers(array []int) []int {
	// Write your code here.

	first := math.MinInt32
	second := math.MinInt32
	third := math.MinInt32

	for i := 0; i < len(array); i++ {
		if array[i] > first {
			third = second
			second = first
			first = array[i]
			continue
		} else if array[i] > second {
			third = second
			second = array[i]
			continue
		} else if array[i] > third {
			third = array[i]
		}
	}
	return []int{third, second, first}
}
