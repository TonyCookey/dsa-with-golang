package main

import "fmt"

func main() {
	xi := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0, 23, 20, 18, 28}
	fmt.Println(selectionSort(xi))
}

func selectionSort(xi []int) []int {

	for i := 0; i < len(xi); i++ {
		var index int
		minValue := xi[i]
		for j := i + 1; j < len(xi); j++ {
			if minValue > xi[j] {
				minValue = xi[j]
				index = j
			}
		}
		if minValue != xi[i] {
			holding := xi[i]
			xi[i] = minValue
			xi[index] = holding
		}
		fmt.Println(xi)
	}
	return xi
}
