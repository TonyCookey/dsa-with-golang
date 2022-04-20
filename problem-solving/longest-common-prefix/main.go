package main

import "strings"

func longestCommonPrefix(strs []string) string {
	prefix := strings.Builder{}

	if len(strs) < 1 || len(strs[0]) < 1 {
		return ""
	}

	for prefixLength := 0; prefixLength < len(strs[0]); prefixLength++ {
		currChar := strs[0][prefixLength]

		for j := 0; j < len(strs); j++ {
			if len(strs[j]) <= prefixLength {
				return prefix.String()
			}
			if currChar != strs[j][prefixLength] {
				return prefix.String()
			}
		}
		prefix.WriteByte(currChar)
	}
	return prefix.String()
}
