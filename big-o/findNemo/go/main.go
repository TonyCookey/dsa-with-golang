package main

import (
	"fmt"
	"time"
)

func main() {
	people := []string{"Nemo", "Nema"}
	start := time.Now()

	for i := 0; i < len(people); i++ {
		if people[i] == "Nemo" {
			fmt.Println("Found NEMO!")
		}
	}
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())

}
