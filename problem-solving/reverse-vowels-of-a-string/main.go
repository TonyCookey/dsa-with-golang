package reverse_vowels_of_a_string

func reverseVowels(s string) string {
	str := []byte(s)

	for i, j := 0, len(str)-1; i < j; {
		if isVowel(str[i]) && isVowel(str[j]) {
			str[i], str[j] = str[j], str[i]
			i += 1
			j -= 1
			continue
		}
		if !isVowel(str[i]) {
			i += 1
		}
		if !isVowel(str[j]) {
			j -= 1
		}
	}
	return string(str)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
