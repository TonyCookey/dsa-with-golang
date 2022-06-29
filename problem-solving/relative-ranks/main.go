package relative_ranks

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	var result []string

	var sortedScore []int
	sortedScore = append(sortedScore, score...)
	scoreMap := make(map[int]int)

	sort.Ints(sortedScore)

	for i := len(sortedScore) - 1; i > 0; i-- {
		scoreMap[sortedScore[i]] = len(sortedScore) - i
	}
	scoreMap[sortedScore[0]] = len(sortedScore)

	for i := 0; i < len(score); i++ {
		var rank string
		if scoreMap[score[i]] == 1 {
			rank = "Gold Medal"
		} else if scoreMap[score[i]] == 2 {
			rank = "Silver Medal"
		} else if scoreMap[score[i]] == 3 {
			rank = "Bronze Medal"
		} else {
			rank = strconv.Itoa(scoreMap[score[i]])
		}
		result = append(result, rank)
	}
	return result
}
