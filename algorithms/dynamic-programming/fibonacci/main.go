package main

import "fmt"

func main() {
	cache := make(map[int]int)
	fibs := FibonacciSequence{
		cache: cache,
	}
	fmt.Println(fibs.fibonacciMemoized(10))
	fmt.Println(fibonacciBottomsUp(10))
}

type FibonacciSequence struct {
	cache map[int]int
}

// fibonacciMemoized - using dynamic programming - memoization and caching
func (f *FibonacciSequence) fibonacciMemoized(n int) int {
	if n < 2 {
		return n
	} else {
		if _, ok := f.cache[n]; ok {
			return f.cache[n]
		} else {
			f.cache[n] = f.fibonacciMemoized(n-1) + f.fibonacciMemoized(n-2)
			return f.cache[n]
		}
	}

}

// fibonacciBottomsUp - using dynamic programming - Bottoms up technique
func fibonacciBottomsUp(n int) int {
	sequence := []int{0, 1}

	for i := 2; i <= n; i++ {
		sequence = append(sequence, (sequence[i-2] + sequence[i-1]))
	}
	return sequence[n]
}
