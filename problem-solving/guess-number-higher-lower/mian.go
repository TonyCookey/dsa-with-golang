package guess_number_higher_lower

func guess(num int) int { return 0 }

func guessNumber(n int) int {
	low := 1
	high := n

	for {
		mid := low + ((high - low) / 2)

		pick := guess(mid)
		if pick == -1 {
			high = mid
		}

		if pick == 1 {
			low = mid + 1
		}

		if pick == 0 {
			return mid
		}
	}
	return 1
}
