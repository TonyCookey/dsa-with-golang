package maximum_area_of_a_piece_of_cake_after_horizontal_and_vertical_cuts

import (
	"math"
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {

	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	maxH := max(h-horizontalCuts[len(horizontalCuts)-1], horizontalCuts[0]-0)

	maxV := max(w-verticalCuts[len(verticalCuts)-1], verticalCuts[0]-0)

	for i := 1; i < max(len(horizontalCuts), len(verticalCuts)); i++ {
		if i < len(horizontalCuts) {
			maxH = max(horizontalCuts[i]-horizontalCuts[i-1], maxH)
		}
		if i < len(verticalCuts) {
			maxV = max(verticalCuts[i]-verticalCuts[i-1], maxV)
		}

	}

	return (maxV * maxH) % (int(math.Pow10(9)) + 7)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
