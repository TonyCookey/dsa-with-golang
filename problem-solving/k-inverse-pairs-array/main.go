package k_inverse_pairs_array

import "math"

func kInversePairs(n int, k int) int {
	curr := make([]int, k+1)
	prev := make([]int, k+1)
	for i := 1; i < n+1; i++ {
		value := 0
		for j := 0; j < k+1; j++ {
			if j == 0 {
				curr[j] = 1
				value += prev[j]
			} else {
				if j >= i {
					value -= prev[j-i]
				}
				value += prev[j]
				curr[j] = value % int(math.Pow(10, 9)+7)
			}
		}
		curr, prev = prev, curr
	}
	return prev[len(prev)-1]
}
