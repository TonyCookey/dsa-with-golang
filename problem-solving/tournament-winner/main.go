package tournament_winner

func TournamentWinner(competitions [][]string, results []int) string {
	scoreboard := make(map[string]int)

	for match, teams := range competitions {
		if results[match] == 1 {
			scoreboard[teams[0]] += 3
		} else {
			scoreboard[teams[1]] += 3
		}
	}
	var max int
	var winner string

	for team, score := range scoreboard {
		if score > max {
			winner = team
			max = score
		}
	}
	return winner
}
