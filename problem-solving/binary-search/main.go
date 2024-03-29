package main

func BinarySearch(array []int, target int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			high = mid - 1
		} else if array[mid] < target {
			low = mid + 1
		}
	}
	return -1
}
