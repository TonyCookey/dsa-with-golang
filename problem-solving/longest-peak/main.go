package main

func LongestPeak(array []int) int {
	left, right := 0, 1
	var peakCount int
	for right <= len(array)-2 {
		if array[right] <= array[right-1] {
			left = right
			right += 1
			continue
		}

		if array[right] > array[right-1] && array[right] > array[right+1] {
			currPeakCount := (right + 1) - left
			for j := right; j <= len(array)-2; j++ {
				if array[j] > array[j+1] {
					currPeakCount++
				} else {
					break
				}
			}
			left = right
			right += 1
			peakCount = max(peakCount, currPeakCount)
			continue
		}

		if array[right] > array[right-1] {
			right++
			continue
		}
	}
	return peakCount
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
