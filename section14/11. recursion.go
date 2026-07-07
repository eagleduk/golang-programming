package main

import (
	"fmt"
)

func main() {

	i := 10

	f := factorial(i)
	fmt.Println("Recursion:", f)

	l := loop(i)
	fmt.Println("Loop:", l)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loop(n int) int {

	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}
