package main

import (
	"fmt"
	"sort"
)

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	// Write your code here.
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)
	var sum int
	for i, j := 0, len(redShirtSpeeds)-1; i < len(redShirtSpeeds) && j >= 0; i, j = i+1, j-1 {
		if fastest {
			sum += max(redShirtSpeeds[i], blueShirtSpeeds[j])
		} else {
			sum += max(redShirtSpeeds[i], blueShirtSpeeds[i])
		}

		fmt.Println(sum)
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
