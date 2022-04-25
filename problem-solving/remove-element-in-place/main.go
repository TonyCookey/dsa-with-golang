package remove_element_in_place

func removeElement(nums []int, val int) int {
	seenValIndex := -1
	var uniqueInt int

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			if seenValIndex == -1 {
				seenValIndex = i
			}
		} else {
			uniqueInt++
			if seenValIndex != -1 {
				nums[i], nums[seenValIndex] = nums[seenValIndex], nums[i]
				seenValIndex++
			}
		}

	}
	return uniqueInt
}
