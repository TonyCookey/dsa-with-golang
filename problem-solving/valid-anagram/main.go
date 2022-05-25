package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	seen := make(map[rune]int)

	for _, char := range s {
		seen[char]++
	}

	for _, char := range t {
		if seen[char] == 0 {
			return false
		}
		seen[char]--
	}
	return true
}
