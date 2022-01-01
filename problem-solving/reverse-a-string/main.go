package main

import (
	"fmt"
	"strings"
)

func main() {
	randomString := "My name is Tony Cookey"
	fmt.Println(reverseStringUsingSplitAfter(randomString, ""))
	fmt.Println(reverseStringUsingSliceOfByte(randomString))
	fmt.Println(reverseStringRecursion(randomString))

	randomSliceOfByte := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(reverseString(randomSliceOfByte))

}

//reverseStringRecursion - reversing string using recursion
func reverseStringRecursion(s string) string {
	if len(s) == 0 {
		return ""
	}
	return reverseStringRecursion(s[1:]) + string(s[0])
}

// reverseStringUsingSplitAfter - reverse string using SplitAfter
func reverseStringUsingSplitAfter(s string, sep string) string {
	// convert string in slice suing strings.SplitAfter()
	// can be used to reverse the words in a string depending on the separator
	//can also be used to reverse every character in the string
	sliceOfString := strings.SplitAfter(s, sep)
	var resultString []string

	for i := len(sliceOfString) - 1; i >= 0; i-- {
		resultString = append(resultString, sliceOfString[i])
	}
	return strings.Join(resultString, sep)

}

//reverseStringUsingSliceOfByte - reverse string using slice of bytes
func reverseStringUsingSliceOfByte(s string) string {
	//convert to a slice of bytes
	//used to reverse every character in the string
	sliceOfByte := []byte(s)
	var resultBytes []byte

	for i := len(sliceOfByte) - 1; i >= 0; i-- {
		resultBytes = append(resultBytes, sliceOfByte[i])
	}
	return string(resultBytes)
}

//reverseString - reverse string (slice of bytes) using in-place algorithm
func reverseString(s []byte) []byte {
	length := len(s)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
