package ugly_number

func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	factors := []int{2, 3, 5}

	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1

}
