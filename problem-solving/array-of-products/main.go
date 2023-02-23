package main

//ArrayOfProducts - 0(N) Time | 0(N) Space
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

//ArrayOfProductsBruteForce - 0(N^2) Time | 0(N) Space
func ArrayOfProductsBruteForce(array []int) []int {
	// Write your code here.
	var result []int
	for i := 0; i < len(array); i++ {
		sum := 1
		for left := i - 1; left >= 0; left-- {
			sum *= array[left]
		}
		for right := i + 1; right < len(array); right++ {
			sum *= array[right]
		}
		result = append(result, sum)
	}

	return result
}
