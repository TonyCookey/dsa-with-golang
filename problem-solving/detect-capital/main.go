package detect_capital

func detectCapitalUse(word string) bool {
	var isFirstCapital bool
	var isCapitalCount int
	var isSmallCount int

	for i, char := range word {

		if i == 0 && isCapital(char) {
			isFirstCapital = true
		}
		if isCapital(char) {
			isCapitalCount++
		} else {
			isSmallCount++
		}
	}
	if isCapitalCount == len(word) || isSmallCount == len(word) || (isFirstCapital && isSmallCount == len(word)-1) {
		return true
	}
	return false

}

func isCapital(char rune) bool {
	if char >= 65 && char <= 90 {
		return true
	}
	return false
}
