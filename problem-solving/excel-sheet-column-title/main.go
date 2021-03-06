package excel_sheet_column_title

import "strings"

func convertToTitle(columnNumber int) string {
	var title strings.Builder
	asciiBase := 64

	for columnNumber > 0 {
		if columnNumber <= 26 {
			title.WriteString(string(columnNumber + asciiBase))
			break
		} else if columnNumber%26 == 0 {
			title.WriteString(string("Z"))
			columnNumber -= 1

		} else {
			title.WriteString(string(columnNumber%26 + asciiBase))
		}
		columnNumber /= 26
	}
	return reverseString(title.String())
}

func reverseString(s string) string {
	var str strings.Builder
	for j := len(s) - 1; j >= 0; j-- {
		str.WriteByte(s[j])
	}
	return str.String()
}
