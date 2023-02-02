package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(133341))
}

// isPalindrome - checks to see if a number is a palindrome(reads the same to and fro - from front and back)
func isPalindrome(x int) bool {
	chars := []rune(strconv.Itoa(x))
	i := 0
	j := len(chars) - 1

	for i < j {
		if chars[i] != chars[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func IsPalindrome(str string) bool {
	start := 0
	end := len(str) - 1
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
