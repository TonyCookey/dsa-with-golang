package main

func MoveElementToEnd(array []int, toMove int) []int {
	left := 0
	right := len(array) - 1

	for left < right {
		for left < right && array[right] == toMove {
			right--
		}
		if array[left] == toMove {
			array[left], array[right] = array[right], array[left]
		}
		left++
	}
	return array
}
