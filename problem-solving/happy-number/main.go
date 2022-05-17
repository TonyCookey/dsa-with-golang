package happy_number

import (
	"math"
)

func isHappy(n int) bool {
	seen := make(map[int]int)
	for {
		if seen[n] > 1 {
			return false
		}
		n = sumDigitsArray(convertToDigitsArray(n))
		seen[n]++
		if n == 1 {
			return true
		}
	}
	return false
}

func sumDigitsArray(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += int(math.Pow(float64(v), 2))
	}
	return sum

}

func convertToDigitsArray(n int) []int {
	var result []int

	for n > 0 {
		result = append(result, n%10)
		n = n / 10
	}
	return result
}
