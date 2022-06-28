package minimum_deletions_to_make_character_freqeuncy_unique

func minDeletions(s string) int {
	seen := make(map[rune]int)

	occurrences := make([]int, len(s)+1)
	var deletions int

	for _, char := range s {
		seen[char]++
	}

	for _, v := range seen {
		if occurrences[v] == 0 {
			occurrences[v] = v
		} else {
			for occurrences[v] != 0 {
				v--
				deletions++
			}
			occurrences[v] = v
		}
	}
	return deletions
}
