package length_of_last_word

func lengthOfLastWord(s string) int {
	var counter int
	var char byte
	for i := len(s) - 1; i >= 0; i-- {
		char = s[i]

		if char != ' ' {
			counter++
		} else {
			if counter > 0 {
				return counter
			} else {
				counter = 0
			}
		}
	}
	return counter

}
