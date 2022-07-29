package find_and_replace_pattern

func findAndReplacePattern(words []string, pattern string) []string {
	var result []string

	for _, word := range words {
		if isPattern(pattern, word) {
			result = append(result, word)
		}
	}
	return result
}

func isPattern(pattern, word string) bool {
	occurences := make(map[rune]byte)
	for i, char := range pattern {
		if val, exists := occurences[char]; exists {
			if val != word[i] {
				return false
			}
		} else {
			occurences[char] = word[i]
		}
	}
	occurences = make(map[rune]byte)
	for i, char := range word {
		if val, exists := occurences[char]; exists {
			if val != pattern[i] {
				return false
			}
		} else {
			occurences[char] = pattern[i]
		}
	}

	return true
}
