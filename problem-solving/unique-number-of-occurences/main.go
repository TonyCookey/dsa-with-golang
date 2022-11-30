package unique_number_of_occurences

import "sort"

//uniqueOccurrences1 - Runs in [NLogN Time , 0(N) Space]
func uniqueOccurrences1(arr []int) bool {
	sort.Ints(arr)

	store := make(map[int]bool)
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			if store[count] {
				return false
			}
			store[count] = true
			count = 1
		} else {
			count++
		}
	}
	if store[count] {
		return false
	}
	return true
}

//uniqueOccurrences -Runs in [0(N) Time, 0(N) Space]
func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)

	for _, val := range arr {
		occurrences[val]++
	}

	checkOccurrences := make(map[int]bool)

	for _, val := range occurrences {
		if checkOccurrences[val] {
			return false
		}
		checkOccurrences[val] = true
	}

	return true
}
