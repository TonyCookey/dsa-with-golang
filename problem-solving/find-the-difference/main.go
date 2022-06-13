package find_the_difference

func findTheDifference(s string, t string) byte {
	seen := make(map[byte]int)

	for i := 0; i < max(len(s), len(t)); i++ {
		if i < len(s) {
			seen[s[i]]++
		}

		if i < len(t) {
			seen[t[i]]++
		}
	}
	for k, v := range seen {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
