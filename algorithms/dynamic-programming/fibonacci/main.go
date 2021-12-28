package main

import "fmt"

func main() {
	cache := make(map[int]int)
	fibs := FibonacciSequence{
		cache: cache,
	}
	fmt.Println(fibs.fibonacciMemoized(10))
}

type FibonacciSequence struct {
	cache map[int]int
}

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
