package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(133341))
}

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
