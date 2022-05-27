package add_digits

func addDigits(num int) int {
	var sum int

	for num > 0 {
		sum += num % 10
		num /= 10
		if num == 0 && sum/10 == 0 {
			return sum
		}

		if num == 0 && sum/10 != 0 {
			num = sum
			sum = 0
		}
	}
	return sum
}
