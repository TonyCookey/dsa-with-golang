package nth_tribonacci_number

func tribonacci(n int) int {
	memo := make(map[int]int)

	var tribHelper func(num int) int

	tribHelper = func(num int) int {
		if num == 0 {
			return num
		}

		if num < 3 {
			return 1
		}

		if val, ok := memo[num]; ok {
			return val
		}

		memo[num] = tribHelper(num-3) + tribHelper(num-2) + tribHelper(num-1)

		return memo[num]
	}
	return tribHelper(n)
}
