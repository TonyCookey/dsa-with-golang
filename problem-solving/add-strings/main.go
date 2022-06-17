package add_strings

import (
	"math"
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	var carry int
	var result strings.Builder

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; {
		var sum, num int
		if i >= 0 {
			num, _ = strconv.Atoi(string(num1[i]))
			sum += num
			i--
		}
		if j >= 0 {
			num, _ = strconv.Atoi(string(num2[j]))
			sum += num
			j--
		}
		sum += carry
		carry = sum / 10
		if sum > 9 {
			sum = sum % 10
		}
		result.WriteString(strconv.Itoa(sum))
	}
	if carry == 1 {
		result.WriteString("1")
	}
	return reverseString(result.String())
}

func reverseString(str string) string {
	var result strings.Builder
	for j := len(str) - 1; j >= 0; j-- {
		result.WriteByte(str[j])
	}
	return result.String()
}

func addStrings2(num1 string, num2 string) string {
	var carry, power int
	var result int

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; {
		var sum, num int
		if i >= 0 {
			num, _ = strconv.Atoi(string(num1[i]))
			sum += num
			i--
		}
		if j >= 0 {
			num, _ = strconv.Atoi(string(num2[j]))
			sum += num
			j--
		}
		sum += carry
		carry = sum / 10
		if sum > 9 {
			sum = sum % 10
		}
		result += sum * int(math.Pow10(power))
		power += 1
	}
	if carry == 1 {
		result += 1 * int(math.Pow10(power))
	}
	return strconv.Itoa(result)
}
