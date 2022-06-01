package power_of_three

import "fmt"

func isPowerOfThree1(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		fmt.Println(n)
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}

func isPowerOfThree2(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
