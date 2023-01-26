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

func BubbleSort(array []int) []int {
	// Write your code here.
	swap := true
	end := len(array) - 1
	for swap {
		swap = false
		for i := 1; i <= end; i++ {
			if array[i] < array[i-1] {
				array[i], array[i-1] = array[i-1], array[i]
				swap = true
			}
		}
		end--
	}
	return array
}
