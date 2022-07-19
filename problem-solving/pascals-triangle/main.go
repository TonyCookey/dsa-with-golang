package pascals_triangle

func generate(numRows int) [][]int {
	result := [][]int{{1}}

	if numRows == 1 {
		return result
	}
	result = append(result, []int{1, 1})

	for i := 3; i <= numRows; i++ {
		row := make([]int, i)
		row[0], row[len(row)-1] = 1, 1

		for j := 1; j < len(row)-1; j++ {
			prevRow := result[len(result)-1]
			row[j] = prevRow[j] + prevRow[j-1]
		}
		result = append(result, row)
	}
	return result
}
