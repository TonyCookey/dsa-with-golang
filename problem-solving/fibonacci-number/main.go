package fibonacci_number

func fib(n int) int {
	seen := make(map[int]int)
	seen[0], seen[1] = 0, 1
	return fibHelper(n, seen)
}

func fibHelper(n int, seen map[int]int) int {
	if _, ok := seen[n]; ok {
		return seen[n]
	}

	seen[n] = fibHelper(n-1, seen) + fibHelper(n-2, seen)
	return seen[n]
}
