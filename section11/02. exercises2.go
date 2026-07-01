package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10, 10)
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	a[5] = 6
	a[6] = 7
	a[7] = 8
	a[8] = 9
	a[9] = 10

	for i, v := range a {
		fmt.Println("[", i, "] = ", v)
	}

	fmt.Printf("%T", a)
}
