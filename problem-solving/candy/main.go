package candy

func candy(ratings []int) int {
	RTL := make([]int, len(ratings))
	LTR := make([]int, len(ratings))

	// for i , _ := range ratings{
	//     RTL[i], LTR[i]  = 1, 1
	// }

	for i := 0; i < len(ratings); i++ {
		LTR[i] = 1
		if i == 0 {
			continue
		}
		if ratings[i] > ratings[i-1] {
			LTR[i] = LTR[i-1] + 1
		}
	}

	for i := len(ratings) - 1; i >= 0; i-- {
		RTL[i] = 1
		if i == len(ratings)-1 {
			continue
		}
		if ratings[i] > ratings[i+1] {
			RTL[i] = RTL[i+1] + 1
		}
	}

	var minNum int

	for i, _ := range ratings {
		minNum += max(LTR[i], RTL[i])
	}
	return minNum

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
