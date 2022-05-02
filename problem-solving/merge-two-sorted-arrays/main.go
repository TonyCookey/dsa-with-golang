package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 7, 8, 9}
	arr2 := []int{4, 6, 8, 11, 20, 34}
	fmt.Println(mergeTwoSortedArrays(arr1, arr2))
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j int
	var iv, jv int
	var result []int

	for i < m || j < n {
		if i < m {
			iv = nums1[i]
		} else {
			iv = intsets.MaxInt
		}
		if j < n {
			jv = nums2[j]
		} else {
			jv = intsets.MaxInt
		}

		if iv <= jv {
			result = append(result, iv)
			i++
		} else {
			result = append(result, jv)
			j++
		}
	}
	for k, v := range result {
		nums1[k] = v
	}
	fmt.Println(result)
}

//mergeTwoSortedArrays - merge two sorted arrays
func mergeTwoSortedArrays(arr1, arr2 []int) []int {
	arr1 = append(arr1, arr2...)
	sort.Ints(arr1)
	return arr1
}
