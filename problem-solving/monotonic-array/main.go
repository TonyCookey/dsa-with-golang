package main

func IsMonotonic(array []int) bool {
	NonIncreasing, NonDecreasing := true, true

	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			NonIncreasing = false
		}

		if array[i] < array[i-1] {
			NonDecreasing = false
		}
	}
	return NonIncreasing || NonDecreasing
}
