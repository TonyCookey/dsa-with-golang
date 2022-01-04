package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(complementBaseTenInteger(5))
}

func complementBaseTenInteger(n int) int {
	nbin := strconv.FormatInt(int64(n), 2)
	var answer string
	for i := 0; i < len(nbin); i++ {
		if nbin[i] == '0' {
			answer = answer + "1"
		} else {
			answer = answer + "0"
		}
	}
	n64, _ := strconv.ParseInt(answer, 2, 64)
	return int(n64)
}
