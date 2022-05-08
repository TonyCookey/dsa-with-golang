package number_of_good_pairs

func numIdenticalPairs(nums []int) int {
	var pairs int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				pairs += 1
			}
		}
	}
	return pairs
}
func numIdenticalPairsMap(nums []int) int {
	var pairs int
	store := make(map[int]int)

	for _, v := range nums {
		pairs += store[v]
		store[v]++
	}

	return pairs
}
