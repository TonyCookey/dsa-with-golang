package word_pattern

import "strings"

func wordPattern(pattern string, s string) bool {
	seenS := make(map[string]byte)
	seenP := make(map[byte]string)

	ss := strings.Split(s, " ")

	if len(ss) != len(pattern) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if seenS[ss[i]] == 0 {
			seenS[ss[i]] = pattern[i]
		}
		if seenS[ss[i]] != 0 && seenS[ss[i]] != pattern[i] {
			return false
		}
		if seenP[pattern[i]] == "" {
			seenP[pattern[i]] = ss[i]
		}
		if seenP[pattern[i]] != " " && seenP[pattern[i]] != ss[i] {
			return false
		}
	}
	return true
}
