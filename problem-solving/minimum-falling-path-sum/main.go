package minimum_falling_path_sum

func minFallingPathSum(matrix [][]int) int {
	for i := len(matrix) - 2; i >= 0; i-- {
		for j := 0; j < len(matrix[i]); j++ {
			var downplus, min_num int
			downplus = matrix[i][j] + matrix[i+1][j]
			if j == 0 {
				min_num = min(downplus, matrix[i][j]+matrix[i+1][j+1])
			} else if j == len(matrix[i])-1 {
				min_num = min(downplus, matrix[i][j]+matrix[i+1][j-1])
			} else {
				min_num = min(downplus, matrix[i][j]+matrix[i+1][j-1])
				min_num = min(min_num, matrix[i][j]+matrix[i+1][j+1])
			}
			matrix[i][j] = min_num
		}
	}
	ans := matrix[0][0]

	for _, val := range matrix[0] {
		ans = min(ans, val)
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
