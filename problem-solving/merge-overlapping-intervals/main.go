package main

import "sort"
import "fmt"

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	var result [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)

	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] < intervals[i-1][1] {
			intervals[i][1] = intervals[i-1][1]
		}

		if intervals[i][0] <= intervals[i-1][1] {
			intervals[i][0] = intervals[i-1][0]
		} else {
			result = append(result, intervals[i-1])
		}
	}
	result = append(result, intervals[len(intervals)-1])

	return result
}
