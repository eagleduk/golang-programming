package main

import (
	"fmt"
)

func main() {
	var a = [5]int{}

	a[0] = 12
	a[1] = 4
	a[2] = 8
	a[3] = 2
	a[4] = 10

	fmt.Println(a)

	for i, v := range a {
		fmt.Println("[", i, "] = ", v)
	}

	fmt.Printf("%T", a)
}
