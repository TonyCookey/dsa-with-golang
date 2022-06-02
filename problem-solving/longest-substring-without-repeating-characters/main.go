package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("tyeriohtyopiuejjskldhksjjskmcnsdkwnjaksjsd"))
}

//lengthOfLongestSubstring - 0(n2)
func lengthOfLongestSubstring(s string) int {
	var longest, count int
	for i := 0; i < len(s); i++ {
		seen := make(map[byte]bool)
		seen[s[i]] = true
		count = 1
		for j := i + 1; j < len(s); j++ {
			if seen[s[j]] {
				break
			} else {
				count += 1
				seen[s[j]] = true
			}
		}
		fmt.Println(count, longest)
		if count > longest {
			longest = count
		}
	}
	return longest
}

//lengthOfLongestSubstringOptimized - 0(n)
func lengthOfLongestSubstringOptimized(s string) int {
	var left, result int
	seenChar := make(map[byte]bool)

	for right := 0; right < len(s); right++ {
		for seenChar[s[right]] {
			seenChar[s[left]] = false
			left += 1
		}
		seenChar[s[right]] = true
		result = max(result, right-left+1)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
