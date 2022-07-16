package out_of_boundary_paths

import "math"

// Memoization
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	const MOD = 1000000007
	var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	memo := make([][][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, maxMove+1)
			for k := 0; k < maxMove+1; k++ {
				memo[i][j][k] = -1
			}
		}
	}
	var DFS func(r, c, steps int) int

	DFS = func(r, c, steps int) int {
		if steps < 0 {
			return 0
		}
		if r >= m || r < 0 || c >= n || c < 0 {
			return 1
		}
		if memo[r][c][steps] != -1 {
			return memo[r][c][steps]
		}
		sum := 0
		for _, d := range dirs {
			sum += DFS(r+d[0], c+d[1], steps-1) % MOD
		}

		memo[r][c][steps] = sum % MOD

		return memo[r][c][steps] % MOD
	}

	return DFS(startRow, startColumn, maxMove)
}

// Dynamic Programming
func findPaths1(m int, n int, N int, i int, j int) int {
	mod := int(math.Pow10(9) + 7)
	dp := make([][][]int, N+1)
	for i := range dp {
		dp[i] = make([][]int, m)
		for I := range dp[i] {
			dp[i][I] = make([]int, n)
		}
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for s := 1; s <= N; s++ {
		for y := 0; y < m; y++ {
			for x := 0; x < n; x++ {
				for _, dir := range dirs {
					ty := dir[0] + y
					tx := dir[1] + x
					if tx < 0 || ty < 0 || tx >= n || ty >= m {
						dp[s][y][x] += 1
					} else {
						dp[s][y][x] = (dp[s][y][x] + dp[s-1][ty][tx]) % mod
					}
				}
			}
		}
	}
	return dp[N][i][j]
}

//TLE
func findPathsMine(m int, n int, maxMove int, startRow int, startColumn int) int {
	var maxPaths int

	var DFS func(r, c, steps int)

	DFS = func(r, c, steps int) {
		if steps > maxMove {
			return
		}

		if r == m || r < 0 || c == n || c < 0 {
			maxPaths++
			return
		}
		steps += 1

		DFS(r+1, c, steps)
		DFS(r-1, c, steps)
		DFS(r, c+1, steps)
		DFS(r, c-1, steps)
		return
	}
	DFS(startRow, startColumn, 0)
	return maxPaths
}
