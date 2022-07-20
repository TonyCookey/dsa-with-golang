package number_of_matching_sequences

func numMatchingSubseq(s string, words []string) int {
	count := 0
	L := len(s)

	wordmap := make(map[string]int)
	for _, word := range words {
		wordmap[word]++
	}

	for word, val := range wordmap {
		w := len(word)
		j := 0
		for i := 0; i < L && j < w; i++ {
			if s[i] == word[j] {
				j++
			}
		}

		if j == w {
			count += val
		}
	}
	return count
}
