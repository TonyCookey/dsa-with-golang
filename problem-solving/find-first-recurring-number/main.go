package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 2}
	fmt.Println(findFirstRecurringNumber(nums))
}

//findFirstRecurringNumber - find first recurring number using hash tables/maps
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
