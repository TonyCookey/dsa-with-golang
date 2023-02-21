package main

func SpiralTraverse(array [][]int) []int {
	var ans []int

	startingCol, endingCol := 0, len(array[0])-1
	startingRow, endingRow := 0, len(array)-1

	for startingRow <= endingRow && startingCol <= endingCol {
		// go right
		for col := startingCol; col <= endingCol; col++ {
			ans = append(ans, array[startingRow][col])
		}
		//go down
		for row := startingRow + 1; row <= endingRow; row++ {
			ans = append(ans, array[row][endingCol])
		}
		//go left
		for col := endingCol - 1; col >= startingCol; col-- {
			// Handle the edge case when there's a single row in the middle of the matrix
			if startingRow == endingRow {
				break
			}
			ans = append(ans, array[endingRow][col])
		}
		//go up
		for row := endingRow - 1; row > startingRow; row-- {
			// Handle the edge case when there's a single column in the middle of the matrix
			if startingCol == endingCol {
				break
			}
			ans = append(ans, array[row][startingCol])
		}
		startingRow++
		startingCol++
		endingCol--
		endingRow--

	}
	return ans
}
