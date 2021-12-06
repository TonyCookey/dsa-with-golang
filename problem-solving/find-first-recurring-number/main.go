package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 2}
	fmt.Println(findFirstRecurringNumber(nums))
}
func findFirstRecurringNumber(nums []int) int {
	uniqueNums := make(map[int]bool)

	for _, num := range nums {
		if uniqueNums[num] {
			return num
		}
		uniqueNums[num] = true
	}
	return 0
}
