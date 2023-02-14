package main

func GenerateDocument(characters string, document string) bool {
	// Write your code here.
	charMap := make(map[rune]int)

	for _, char := range characters {
		charMap[char]++
	}
	for _, char := range document {
		charMap[char]--
	}
	for _, count := range charMap {
		if count < 0 {
			return false
		}
	}
	return true
}

func GenerateDocument2(characters string, document string) bool {
	// Write your code here.
	charMap := make(map[rune]int)

	for _, char := range characters {
		charMap[char]++
	}
	for _, char := range document {

		charMap[char]--
		if charMap[char] < 0 {
			return false
		}
	}
	return true
}
