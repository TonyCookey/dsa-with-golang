package excel_sheet_column_number

import "math"

func titleToNumber(columnTitle string) int {
	index := float64(len(columnTitle) - 1)
	var finalColNum float64

	for _, char := range columnTitle {
		num := float64(char - 64)
		finalColNum += num * math.Pow(26, index)
		index--
	}
	return int(finalColNum)
}
