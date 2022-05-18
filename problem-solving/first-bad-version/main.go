package first_bad_version

import "fmt"

func isBadVersion(version int) bool {
	return false
}
func firstBadVersionLinear(n int) int {
	var isBad bool
	for i := n; n > 0; i-- {
		if !isBadVersion(i) && isBad {
			return i + 1
		}
		if isBadVersion(i) {
			isBad = true
		} else {
			isBad = false
		}
	}
	return n
}

func firstBadVersionBST(n int) int {
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2
		fmt.Println(mid)
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
