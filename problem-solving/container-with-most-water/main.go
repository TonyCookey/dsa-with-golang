package container_with_most_water

func maxArea(height []int) int {
	var highestArea int
	left := 0
	right := len(height) - 1

	for left < right {
		area := (right - left) * min(height[left], height[right])
		highestArea = max(highestArea, area)
		if height[left] == min(height[left], height[right]) {
			left += 1
		} else {
			right -= 1
		}
	}
	return highestArea
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
