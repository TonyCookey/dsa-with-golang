package main

import "fmt"

func main() {
	xi := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0, 23, 20, 18, 28}
	fmt.Println(insertionSort(xi))
}

func InsertionSort(array []int) []int {

	for i := 1; i < len(array); i++ {
		index := i
		for j := i - 1; j >= 0; j-- {
			if array[index] < array[j] {
				array[index], array[j] = array[j], array[index]
			}
			index--
		}
	}
	return array
}

//insertionSort - insertion sort algorithm
func insertionSort(xi []int) []int {
	for i := 0; i < len(xi)-1; i++ {
		holding := xi[i]
		if xi[i] > xi[i+1] {
			xi[i] = xi[i+1]
			xi[i+1] = holding
			for j := i; j > 0; j-- {
				if xi[j] < xi[j-1] {
					holding = xi[j]
					xi[j] = xi[j-1]
					xi[j-1] = holding
				}
			}
		}
	}
	return xi
}
