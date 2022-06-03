package reverse_integer

import "math"

func reverse(x int) int {
	var isNegative bool
	if x < 0 {
		isNegative = true
		x = -x
	}
	var xs []int

	for x > 0 {
		xs = append(xs, x%10)
		x /= 10
	}
	var sum int

	for i, j := 0, len(xs)-1; i < len(xs); i, j = i+1, j-1 {
		sum += xs[i] * int(math.Pow10(j))
		// greater than int32
		if sum > 2147483647 {
			return 0
		}
	}

	if isNegative {
		return -sum
	}
	return sum

}
