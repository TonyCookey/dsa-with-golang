package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	mapping := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if mapping[s[i]] == 0 {
			mapping[s[i]] = t[i]

		}
		if mapping[s[i]] != 0 && mapping[s[i]] != t[i] {
			return false
		}
	}

	mapping = make(map[byte]byte)
	for i := 0; i < len(t); i++ {
		if mapping[t[i]] == 0 {
			mapping[t[i]] = s[i]

		}
		if mapping[t[i]] != 0 && mapping[t[i]] != s[i] {
			return false
		}
	}
	return true
}
