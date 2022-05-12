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
