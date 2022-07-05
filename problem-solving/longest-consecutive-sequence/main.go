package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	seen := make(map[int]bool)
	var maxV, minV int

	for _, v := range nums {
		seen[v] = true
		maxV = max(maxV, v)
		minV = min(minV, v)
	}
	count := 1
	maxCount := 0
	cant := 0
	for i := minV + 1; i <= maxV; i++ {
		if seen[i] && seen[i-1] {
			count++
			maxCount = max(maxCount, count)
		} else {
			maxCount = max(maxCount, count)
			count = 1
			cant++
			if maxV-i > i-minV && cant > maxCount*len(nums) {
				break
			}
		}
	}
	return max(maxCount, count)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
