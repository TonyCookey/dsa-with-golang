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

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	seen := make(map[byte]int)

	for i, _ := range s {
		seen[s[i]]++
		seen[t[i]]--
	}
	for _, char := range seen {
		if char != 0 {
			return false
		}
	}
	return true
}
