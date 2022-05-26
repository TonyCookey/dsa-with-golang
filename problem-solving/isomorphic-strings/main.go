package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	forwardMapping := make(map[byte]byte)
	backwardMapping := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if forwardMapping[s[i]] == 0 {
			forwardMapping[s[i]] = t[i]
		}
		if backwardMapping[t[i]] == 0 {
			backwardMapping[t[i]] = s[i]
		}
		if forwardMapping[s[i]] != 0 && forwardMapping[s[i]] != t[i] {
			return false
		}
		if backwardMapping[t[i]] != 0 && backwardMapping[t[i]] != s[i] {
			return false
		}
	}
	return true
}
