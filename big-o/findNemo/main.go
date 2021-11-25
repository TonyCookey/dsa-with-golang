package main

import "fmt"

func main() {
	people := []string{"Nemo", "Nema"}
	for i := 0; i < len(people); i++ {
		if people[i] == "Nemo" {
			fmt.Println("Found NEMO!")
		}
	}
}
