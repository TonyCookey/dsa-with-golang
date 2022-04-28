package sqrt_x

func mySqrt(x int) int {
	low, high := 0, x
	var mid int
	for low <= high {
		mid = low + ((high - low) / 2)
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if high < low {
		return high
	}
	return low
}
