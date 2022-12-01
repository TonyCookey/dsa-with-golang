package determine_if_string_halves_are_alike

//halvesAreAlike - [0(N) Time, 0(N) Space]
func halvesAreAlike(s string) bool {
	a := s[:len(s)/2]
	b := s[len(s)/2:]

	var aCount int
	var bCount int

	for _, char := range a {
		if isVowel1(char) {
			aCount++
		}
	}
	for _, char := range b {
		if isVowel1(char) {
			bCount++
		}
	}
	return aCount == bCount
}
func isVowel1(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}

//halvesAreAlike2 - [0(N) Time, 0(1) Space]
func halvesAreAlike2(s string) bool {
	var aCount int
	var bCount int

	for i := 0; i < len(s)/2; i++ {
		if isVowel(s[i]) {
			aCount++
		}
	}
	for i := len(s) / 2; i < len(s); i++ {
		if isVowel(s[i]) {
			bCount++
		}
	}
	return aCount == bCount
}

func isVowel(r uint8) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}
