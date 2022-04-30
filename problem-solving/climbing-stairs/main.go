package climbing_stairs

func climbingStairsMaths(n int) int {
	a, b := 1, 1
	for n > 0 {
		b += a
		a = b - a
		n--
	}
	return a
}
func climbStairs(n int) int {
	if n <= 3 {
		return n
	}

	hashmap := make(map[int]int)

	return getDistinctPaths(n-2, &hashmap) + getDistinctPaths(n-1, &hashmap)
}

func getDistinctPaths(n int, hashmap *map[int]int) int {
	if (*hashmap)[n] > 0 {
		return (*hashmap)[n]
	}

	if n <= 0 {
		if n == 0 {
			return 1
		}
		return 0
	}

	(*hashmap)[n] = getDistinctPaths(n-2, hashmap) + getDistinctPaths(n-1, hashmap)
	return (*hashmap)[n]
}
