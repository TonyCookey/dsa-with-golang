package add_binary

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	var result strings.Builder
	var carry int
	var sum int
	var ai, bj int
	for i, j := len(a)-1, len(b)-1; i > -1 || j > -1; i, j = i-1, j-1 {

		if i > -1 {
			ai = convertStringToInt(a[i])
			sum += ai
		}
		if j > -1 {
			bj = convertStringToInt(b[j])
			sum += bj
		}
		sum = carry + sum
		carry = sum / 2
		result.WriteString(strconv.Itoa(sum % 2))
		sum = 0
	}
	if carry == 1 {
		result.WriteString("1")
	}
	fmt.Println(result.String())
	return reverseString(result.String())
}

func convertStringToInt(i byte) int {
	if i == 49 {
		return 1
	} else {
		return 0
	}
}
func reverseString(str string) string {
	buf := []rune(str)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
