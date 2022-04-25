package implement_str_str

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	var chars string
	needleLen := len(needle)
	for i := 0; i <= len(haystack)-needleLen; i++ {
		chars = haystack[i : i+needleLen]
		if chars == needle {
			return i
		}
	}
	return -1

}
