package word_subsets

func wordSubsets(A []string, B []string) []string {
	ans, tmp := []string{}, make([]int, 26)
	for _, b := range B {
		count := helper(b)
		for i := 0; i < 26; i++ {
			if count[i] > tmp[i] {
				tmp[i] = count[i]
			}
		}
	}
	for _, a := range A {
		count, b := helper(a), false
		for i := 0; i < 26; i++ {
			if count[i] < tmp[i] {
				break
			}
			if i == 25 {
				b = true
			}
		}
		if b {
			ans = append(ans, a)
		}
	}
	return ans
}

func helper(s string) []int {
	tmp := make([]int, 26)
	for _, c := range s {
		tmp[c-'a']++
	}
	return tmp
}
