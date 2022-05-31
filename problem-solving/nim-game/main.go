package nim_game

func canWinNim(n int) bool {
	if n > 1 && n < 4 {
		return true
	}
	return n%4 != 0
}
