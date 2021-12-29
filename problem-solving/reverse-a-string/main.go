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

}
func reverseStringRecursion(s string) string {
	if len(s) == 0 {
		return ""
	}
	return reverseStringRecursion(s[1:]) + string(s[0])
}
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
