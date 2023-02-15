package main

func FirstNonRepeatingCharacter(str string) int {
	charMap := make(map[rune]int)

	for _, char := range str {
		charMap[char]++
	}
	for index, char := range str {
		if charMap[char] == 1 {
			return index
		}
	}

	return -1
}
