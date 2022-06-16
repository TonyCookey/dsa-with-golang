package longest_palindrome

func longestPalindrome(s string) int {
	seen := make(map[rune]int)
	var result int
	var isOne bool
	for _, char := range s {
		seen[char]++
	}

	for _, v := range seen {
		if v%2 == 0 {
			result += v
		}
		if v%2 != 0 && v > 2 {
			if isOne {
				result += v - 1
				continue
			}
			result += v
			isOne = true
		}
		if v == 1 && !isOne {
			result += v
			isOne = true
		}
	}

	return result
}
