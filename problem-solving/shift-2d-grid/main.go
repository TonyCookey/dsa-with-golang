package main

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	var result [][]int
	var last int
	for i := 0; i < k; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				last = grid[m-1][n-1]
			} else {
				last = grid[j-1][n-1]
			}
			row := append([]int{last}, grid[j][0:n-1]...)

			result = append(result, row)
		}
		grid = result
		result = [][]int{}
	}
	return grid
}
