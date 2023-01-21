package main

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.

func ProductSum(array SpecialArray) int {
	// Write your code here.

	var recursiveProductSum func(array SpecialArray, depth int) int
	recursiveProductSum = func(array SpecialArray, depth int) int {
		var innerSum int
		for _, val := range array {
			if cast, ok := val.(SpecialArray); ok {
				innerSum += recursiveProductSum(cast, depth+1)

			} else if cast, ok := val.(int); ok {
				innerSum += cast
			}
		}
		return depth * innerSum
	}
	return recursiveProductSum(array, 1)
}
