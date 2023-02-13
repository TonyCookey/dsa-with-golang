package main

import "strings"
import "strconv"

func RunLengthEncoding(str string) string {
	// Write your code here.
	var sb strings.Builder
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] != str[i-1] || (str[i] == str[i-1] && count > 8) {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteByte(str[i-1])
			count = 1
		} else {
			count++
		}
	}
	sb.WriteString(strconv.Itoa(count))
	sb.WriteByte(str[len(str)-1])
	return sb.String()
}
