package main

func ArrayOfProducts(array []int) []int {
	var result []int
	cumulativeSum := 1

	backwardSum := make([]int, len(array))

	backwardSum[len(array)-1] = array[len(array)-1]

	currSum := 1
	for i := len(array) - 1; i >= 0; i-- {
		backwardSum[i] = array[i] * currSum
		currSum = backwardSum[i]
	}

	for idx := 0; idx < len(array)-1; idx++ {
		sum := cumulativeSum * backwardSum[idx+1]
		result = append(result, sum)
		cumulativeSum *= array[idx]
	}
	result = append(result, cumulativeSum)
	return result
}
