package max_area_of_island

func maxAreaOfIsland(grid [][]int) int {
	visited := make(map[[2]int]bool)
	var maxArea int
	var DFS func(m, n int) int

	rows, cols := len(grid), len(grid[0])

	DFS = func(r, c int) int {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
			return 0
		}
		if _, exists := visited[[2]int{r, c}]; exists {
			return 0
		}
		visited[[2]int{r, c}] = true

		return 1 + DFS(r+1, c) + DFS(r-1, c) + DFS(r, c-1) + DFS(r, c+1)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			maxArea = max(maxArea, DFS(i, j))
		}
	}

	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
