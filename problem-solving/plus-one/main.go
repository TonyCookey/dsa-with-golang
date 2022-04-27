package plus_one

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] > 8 {
			if i-1 > -1 {
				digits[i] = 0
				if digits[i-1] > 8 {
					digits[i-1]++
				}

			} else {
				digits[i] = 1
				digits = append(digits, 0)
				break
			}
		} else {
			digits[i]++
			break
		}
	}
	return digits
}
