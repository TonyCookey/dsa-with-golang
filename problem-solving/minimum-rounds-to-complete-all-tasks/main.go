package minimum_rounds_to_complete_all_tasks

func minimumRounds(tasks []int) int {
	var rounds int
	taskCount := make(map[int]int)
	for _, val := range tasks {
		taskCount[val]++
	}
	for _, val := range taskCount {
		if val == 1 {
			return -1
		}
		rounds += val / 3

		if val%3 == 1 || val%3 == 2 {
			rounds += 1
		}
	}
	return rounds
}
