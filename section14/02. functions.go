package main

import (
	"fmt"
)

func main() {
	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("total is ", sum)
}

func sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
