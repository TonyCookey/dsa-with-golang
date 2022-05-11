package pascal_triangle_II

func getRow(rowIndex int) []int {
	result := [][]int{{1}, {1, 1}}

	for i := 3; i <= rowIndex+1; i++ {
		row := make([]int, i)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = result[len(result)-1][j-1] + result[len(result)-1][j]
		}
		result = append(result, row)
	}
	return result[rowIndex]
}
