package main

import "fmt"

func main() {
	sampleArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(sampleArray); i++ {
		for j := 0; j < len(sampleArray); j++ {
			fmt.Println("Sample Array:", sampleArray[i], " Item: ", sampleArray[j])
		}
	}
}
