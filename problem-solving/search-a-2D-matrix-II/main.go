package search_a_2D_matrix_II

func searchMatrix(matrix [][]int, target int) bool {
	var hasFound bool
	min := matrix[0][0] - 1
	var DFS func(r, c int)
	DFS = func(r, c int) {
		if hasFound {
			return
		}
		if r == len(matrix) || c == len(matrix[r]) {
			return
		}
		if matrix[r][c] == min {
			return
		}
		if matrix[r][c] == target {
			hasFound = true
		}
		matrix[r][c] = min
		DFS(r+1, c)
		DFS(r, c+1)
	}
	DFS(0, 0)

	return hasFound
}
