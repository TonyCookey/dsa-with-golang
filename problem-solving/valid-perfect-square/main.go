package valid_perfect_square

func isPerfectSquare(num int) bool {
	n := num
	min := 0
	max := num

	for n > 0 {
		n /= 10
		if n*n > num {
			max = n
		}
		if n*n < num {
			min = n
			break
		}
	}

	for ; min <= max; min++ {
		if min*min == num {
			return true
		}
		if min*min > num {
			break
		}
	}

	return false
}
