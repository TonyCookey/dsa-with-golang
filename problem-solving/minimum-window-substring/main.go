package minimum_window_substring

import "golang.org/x/tools/container/intsets"

func minWindow(s string, t string) string {
	if len(t) > len(s) || len(t) == 0 {
		return ""
	}
	haveMap, needMap := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(t); i++ {
		needMap[t[i]]++
	}
	result, resultLen := []int{0, 0}, intsets.MaxInt

	have, need := 0, len(needMap)
	var l int
	for r := 0; r < len(s); r++ {
		haveMap[s[r]]++
		if needMap[s[r]] > 0 && haveMap[s[r]] == needMap[s[r]] {
			have += 1
		}
		for have == need {
			if r-l+1 < resultLen {
				result = []int{l, r}
				resultLen = r - l + 1
			}
			haveMap[s[l]]--
			if needMap[s[l]] > 0 && haveMap[s[l]] < needMap[s[l]] {
				have -= 1
			}
			l++
		}
	}
	if resultLen == intsets.MaxInt {
		return ""
	} else {
		return s[result[0] : result[1]+1]
	}

}
