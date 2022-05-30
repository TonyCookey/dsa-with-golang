package move_zeroes

func moveZeroes(nums []int) {
	lastSeen := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && lastSeen == -1 {
			lastSeen = i
		}

		if nums[i] != 0 && lastSeen != -1 {
			nums[lastSeen] = nums[i]
			nums[i] = 0
			lastSeen += 1
		}
	}
}
