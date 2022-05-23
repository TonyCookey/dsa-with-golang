package power_of_two

import "math"

func isPowerOfTwo(n int) bool {
	var power float64
	x := n
	for n >= 0 {
		if n == 0 {
			return false
		}
		if math.Pow(2, power) == float64(x) {
			return true
		}
		n /= 2
		power += 1
	}
	return false
}
