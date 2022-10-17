package check_if_sentence_is_pangram

func checkIfPangram(sentence string) bool {
	exists := make([]bool, 26, 26)

	for _, char := range sentence {
		exists[char-'a'] = true
	}

	for _, val := range exists {
		if !val {
			return false
		}
	}
	return true
}
