package tournament_winner

func TournamentWinner1(competitions [][]string, results []int) string {
	// Write your code here.
	scoreboard := make(map[string]int)
	var winner string
	scoreboard[winner] = 0

	for match, teams := range competitions {
		var winningTeam string

		winningTeam = teams[1]

		if results[match] == 1 {
			winningTeam = teams[0]
		}
		scoreboard[winningTeam] += 3
		if scoreboard[winningTeam] > scoreboard[winner] {
			winner = winningTeam
		}
	}

	return winner
}
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
