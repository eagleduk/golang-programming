package main

import (
	"fmt"
)

func main() {
	s, c := foo(10)

	fmt.Println(s)
	fmt.Println(c)
}

func foo(c int) (string, int) {
	result := ""
	a := "A"
	b := c

	for i := 0; i < b; i++ {
		result = result + " " + a
	}
	return result, b
}
