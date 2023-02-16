package main

import "strings"

func Semordnilap(words []string) [][]string {
	ans := [][]string{}

	reversedStrMap := make(map[string]bool)

	for _, word := range words {
		revWord := reverseString(word)
		if !reversedStrMap[word] {
			reversedStrMap[revWord] = true
		} else {
			ans = append(ans, []string{word, revWord})
		}
	}
	return ans
}

func reverseString(str string) string {
	var sb strings.Builder

	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte(str[i])
	}
	return sb.String()
}
