package number_of_segments_in_a_string

func countSegments(s string) int {
	var result []string
	var hasSeenChar bool
	for i, j := 0, 0; j < len(s); j++ {
		if s[j] != ' ' {
			hasSeenChar = true
		}
		if (s[j] == ' ' || j == len(s)-1) && hasSeenChar {
			result = append(result, s[i:j])
			i = j
			hasSeenChar = false
		}

	}
	return len(result)
}
