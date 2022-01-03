package main

import "fmt"

func main() {
	//fmt.Println(findJudge(4, [][]int{{1, 3}, {1,4}, {2, 3}, {2,4}, {4,3}}))
	//fmt.Println(findJudge(1, [][]int{}))
	//fmt.Println(findJudge(3, [][]int{{1,2},{2,3}}))
	//[[1,2],[1,3],[2,1],[2,3],[1,4],[4,3],[4,1]]
	fmt.Println(findJudge2(4, [][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {1, 4}, {4, 3}, {4, 1}}))
	//fmt.Println(findJudge(3, [][]int{{1, 3}, {2,3}, {3, 1}}))
	//fmt.Println(findJudge(4, [][]int{{1,2},{3,2},{1,3},{4,1},{3,1},{2,1},{2,3},{4,3},{4,2},{3,4},{2,4}}))
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
