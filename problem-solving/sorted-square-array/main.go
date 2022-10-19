package sorted_square_array

import "sort"

func SortedSquaredArray(array []int) []int {

	for i := 0; i < len(array); i++ {
		array[i] = array[i] * array[i]
	}

	sort.Ints(array)
	return array
}
