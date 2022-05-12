package valid_palindrome

import (
	"fmt"
	"strings"
)

//isPalindrome - Time Complexity - 0(n+m)
func isPalindrome(s string) bool {
	var result strings.Builder
	s = strings.ToLower(s)
	for _, char := range s {
		fmt.Println(char)
		if (char >= 97 && char <= 122) || (char >= 48 && char <= 57) {
			result.WriteRune(char)
		}
	}
	palindrome := result.String()
	fmt.Println(palindrome)
	for i, j := 0, len(palindrome)-1; i < j; i, j = i+1, j-1 {
		if palindrome[i] != palindrome[j] {
			return false
		}
	}
	return true
}

//isPalindromeCleaner - Time Complexity - 0(n)
func isPalindromeCleaner(s string) bool {
	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i < j; {
		left := s[i]
		right := s[j]

		if !isAlphanumeric(left) {
			i++
			continue
		}
		if !isAlphanumeric(right) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 97 && c <= 122) || (c >= 48 && c <= 57)
}
