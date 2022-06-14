package is_subsequence

func isSubsequence(s string, t string) bool {
	var count int
	for i, j := 0, 0; i < len(s) && j < len(t); {
		if s[i] == t[j] {
			count += 1
			j += 1
			i += 1
			continue
		}
		j += 1
	}
	return count == len(s)
}
