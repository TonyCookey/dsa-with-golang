package main

import (
	"golang.org/x/tools/container/intsets"
	"sort"
)

// SmallestDifferenceBruteForce  - 0(NM) Time | 0(1) Space
func SmallestDifferenceBruteForce(array1, array2 []int) []int {
	minDiff := intsets.MaxInt
	var ans []int
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if abs(array1[i]-array2[j]) < minDiff {
				minDiff = abs(array1[i] - array2[j])
				ans = []int{array1[i], array2[j]}
			}
		}
	}
	return ans
}

//SmallestDifferenceOptimal - 0(NLogN + MLogM) Time | 0(1) Space
func SmallestDifferenceOptimal(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)

	minDiff := intsets.MaxInt
	var ans []int

	for i, j := 0, 0; i < len(array1) && j < len(array2); {
		diff := abs(array1[i] - array2[j])
		if diff < minDiff {
			minDiff = diff
			ans = []int{array1[i], array2[j]}
		}

		if array1[i] < array2[j] {
			i++
		} else {
			j++
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
