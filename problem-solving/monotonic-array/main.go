package main

func IsMonotonic(array []int) bool {
	NonIncreasing, NonDecreasing := true, true

	for i := 1; i < len(array); i++ {
		if NonIncreasing && array[i] > array[i-1] {
			NonIncreasing = false
		}

		if NonDecreasing && array[i] < array[i-1] {
			NonDecreasing = false
		}
	}
	if NonIncreasing || NonDecreasing {
		return true
	}
	return false
}
