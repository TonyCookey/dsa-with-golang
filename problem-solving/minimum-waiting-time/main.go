package minimum_waiting_time

import "sort"

// 0(NlogN) Time | 0(1) Space
func MinimumWaitingTime(queries []int) int {
	// Write your code here
	sort.Ints(queries)
	waitingTime := 0
	sum := 0
	for i := 1; i < len(queries); i++ {
		waitingTime += queries[i-1]
		sum += waitingTime
	}
	return sum
}

// 0(NlogN) Time | 0(1) Space
func MinimumWaitingTimeOptimal(queries []int) int {
	// Write your code here
	sort.Ints(queries)
	waitingTime := 0

	for idx, duration := range queries {
		queriesLeft := len(queries) - (idx + 1)
		waitingTime += duration * queriesLeft
	}
	return waitingTime
}
