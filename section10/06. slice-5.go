package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 11)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 123)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 1234)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
