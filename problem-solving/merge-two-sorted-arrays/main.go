package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 7, 8, 9}
	arr2 := []int{4, 6, 8, 11, 20, 34}
	fmt.Println(mergeTwoSortedArrays(arr1, arr2))
}

//mergeTwoSortedArrays - merge two sorted arrays
func mergeTwoSortedArrays(arr1, arr2 []int) []int {
	arr1 = append(arr1, arr2...)
	sort.Ints(arr1)
	return arr1
}
