package main

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	// Write your code here.

	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)

	var redShirtBack, blueShirtBack int

	for i := 0; i < len(redShirtHeights); i++ {
		if redShirtHeights[i] > blueShirtHeights[i] {
			redShirtBack++
		} else if redShirtHeights[i] < blueShirtHeights[i] {
			blueShirtBack++
		}
	}
	if redShirtBack == len(redShirtHeights) || blueShirtBack == len(redShirtHeights) {
		return true
	}
	return false
}
