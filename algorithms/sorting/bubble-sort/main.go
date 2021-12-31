package main

import "fmt"

func main() {
	xi := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0, 23, 20, 18, 28}
	fmt.Println(bubbleSort(xi))
}

//bubbleSort - bubble sort algorithm
func bubbleSort(xi []int) []int {
	length := len(xi)

	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if xi[j] > xi[j+1] {
				current := xi[j]
				xi[j] = xi[j+1]
				xi[j+1] = current
			}
		}
	}
	return xi
}
