package gas_station

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	tank, currTank := 0, 0
	start := 0

	for i := 0; i < n; i++ {
		tank += gas[i] - cost[i]
		currTank += gas[i] - cost[i]
		if currTank < 0 {
			start = i + 1
			currTank = 0
		}
	}
	if tank < 0 {
		start = -1
	}
	return start
}
