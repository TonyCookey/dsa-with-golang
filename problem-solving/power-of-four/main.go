package power_of_four

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}
