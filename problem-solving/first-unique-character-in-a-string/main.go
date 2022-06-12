package first_unique_character_in_a_string

func firstUniqChar(s string) int {
	seen := make(map[rune]int)

	for _, v := range s {
		seen[v]++
	}

	for i, v := range s {
		if seen[v] == 1 {
			return i
		}
	}
	return -1
}
