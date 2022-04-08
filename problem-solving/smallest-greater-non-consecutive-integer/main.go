package main

// you can also use imports, for example:
// import "fmt"
import (
	"fmt"
	"math"
)

// you can write to stdout for debugging purposes, e.g.
func main() {
	fmt.Println(Solution(9999))
}
func Solution(N int) int {
	// write your code in Go 1.4
	digitArray := isConsecutiveDigit(N + 1)
	for {
		x := true
		for i := 0; i < len(digitArray)-1; i++ {
			if digitArray[i] > 9 {
				if i != 0 {

					digitArray[i] -= 9
					digitArray[i-1]++
				}
			}
			if digitArray[i] == digitArray[i+1] {
				digitArray[i+1]++
				if digitArray[i+1] > 9 {
					digitArray[i+1] -= 9
					digitArray[i]++
				}
				digitArray = changeArrayItemsToZeroFromIndex(digitArray, i+1)
				x = false
				break
			}
		}
		if x == true {
			break
		}

	}
	return mergeDigitArrayIntoInteger(digitArray)
}

func isConsecutiveDigit(N int) []int {
	digits := []int{}
	for N > 0 {
		digits = append(digits, N%10)

		N /= 10

	}
	return reverseSlice(digits)
}
func mergeDigitArrayIntoInteger(digits []int) int {
	ans := 0
	count := len(digits) - 1

	for i := 0; i < len(digits); i++ {
		ans += int(math.Pow(10, float64(count))) * digits[i]
		count--
	}
	return ans
}
func changeArrayItemsToZeroFromIndex(digitArray []int, index int) []int {
	for i := index + 1; i < len(digitArray); i++ {
		digitArray[i] = 0
	}
	return digitArray

}
func reverseSlice(slice []int) []int {
	for i := 0; i < len(slice)/2; i++ {
		a := slice[i]
		slice[i] = slice[(len(slice)-1)-i]
		slice[(len(slice)-1)-i] = a
	}
	return slice
}
