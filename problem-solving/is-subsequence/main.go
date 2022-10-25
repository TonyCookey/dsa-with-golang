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

func IsValidSubsequence(array []int, sequence []int) bool {
	var matches int
	for i := 0; i < len(sequence); i++ {
		for j := 0; i < len(sequence) && j < len(array); {
			if sequence[i] == array[j] {
				i++
				j++
				matches++
				continue
			} else {
				j++
			}
		}
	}
	return matches == len(sequence)
}
