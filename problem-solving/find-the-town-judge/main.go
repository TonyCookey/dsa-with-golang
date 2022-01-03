package main

import "fmt"

func main() {
	//fmt.Println(findJudge(4, [][]int{{1, 3}, {1,4}, {2, 3}, {2,4}, {4,3}}))
	//fmt.Println(findJudge(1, [][]int{}))
	//fmt.Println(findJudge(3, [][]int{{1,2},{2,3}}))
	fmt.Println(findJudge2(3, [][]int{{1, 2}, {2, 3}}))
	//[[1,2],[1,3],[2,1],[2,3],[1,4],[4,3],[4,1]]
	fmt.Println(findJudge(4, [][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {1, 4}, {4, 3}, {4, 1}}))
	//fmt.Println(findJudge(3, [][]int{{1, 3}, {2,3}, {3, 1}}))
	//fmt.Println(findJudge(4, [][]int{{1,2},{3,2},{1,3},{4,1},{3,1},{2,1},{2,3},{4,3},{4,2},{3,4},{2,4}}))
}
func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n > 1 {
		return -1
	}
	if len(trust) == 0 && n == 1 {
		return n
	}
	//townJudge := trust[0][1]
	voterMap := map[int]int{trust[0][0]: 1}
	voteeMap := map[int]int{trust[0][1]: 1}

	for i := 1; i < len(trust); i++ {
		voter := trust[i][0]
		votee := trust[i][1]
		//if count, ok := voteeMap[townJudge]; ok{
		//	if count != 0 {
		//		townJudge = voter
		//	}
		//}else{
		//	townJudge = votee
		//}
		if _, ok := voterMap[voter]; ok {
			voterMap[voter]++
		} else {
			voterMap[voter] = 1
		}
		if _, ok := voteeMap[votee]; ok {
			voteeMap[votee]++
		} else {
			voteeMap[votee] = 1
		}
	}
	for key := range voteeMap {
		if voteeMap[key] == n-1 && voterMap[key] == 0 {
			return key
		}
	}
	fmt.Println(voterMap, voteeMap)
	//if voterMap[townJudge] == 0 && voteeMap[townJudge] == n - 1 {
	//	return townJudge
	//}
	return -1
}
func findJudge2(n int, trust [][]int) int {
	townJudge := trust[0][1]
	trustMap := map[int]int{townJudge: 1}

	if n <= 2 {
		return townJudge
	}

	for i := 1; i < len(trust); i++ {
		if trust[i][0] == townJudge {
			trustMap[townJudge]--
		}
		current := trust[i][1]
		if _, ok := trustMap[current]; ok {
			trustMap[current]++
			if trustMap[townJudge] < trustMap[current] {
				townJudge = current
			}
		} else {
			trustMap[current] = 1
		}
	}
	fmt.Println(trustMap)
	if trustMap[townJudge] == n-1 {

		return townJudge
	}
	return -1
}
