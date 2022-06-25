package max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) int {
	var onecount, maxOnesCount int

	for _, v := range nums {
		if v == 1 {
			onecount++
		}
		if v == 0 {
			onecount = 0
		}
		if onecount > maxOnesCount {
			maxOnesCount = onecount
		}
	}
	return maxOnesCount
}
