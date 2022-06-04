package longest_palindromic_substring

func longestPalindrome(s string) string {
	var longest int
	var result string

	for i := 0; i < len(s); i++ {

		for j, k := i, i; j >= 0 && k < len(s) && s[j] == s[k]; j, k = j-1, k+1 {
			if k-j+1 > longest {
				result = s[j : k+1]
				longest = k - j + 1
			}
		}

		for j, k := i, i+1; j >= 0 && k < len(s) && s[j] == s[k]; j, k = j-1, k+1 {
			if k-j+1 > longest {
				result = s[j : k+1]
				longest = k - j + 1
			}
		}

	}

	return result
}
