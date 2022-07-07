package interleaving_strings

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	var DFS func(i, j int) bool
	dp := make(map[[2]int]bool)
	DFS = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}
		if v, ok := dp[[2]int{i, j}]; ok {
			return v
		}
		if i < len(s1) && s1[i] == s3[i+j] && DFS(i+1, j) {
			return true
		}
		if j < len(s2) && s2[j] == s3[i+j] && DFS(i, j+1) {
			return true
		}
		dp[[2]int{i, j}] = false

		return false
	}
	return DFS(0, 0)
}
