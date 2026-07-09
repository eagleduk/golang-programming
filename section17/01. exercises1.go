package main

import (
	"fmt"
)

func main() {
	a := 99

	fmt.Printf("a: %v[%T]\n", a, a)
	fmt.Printf("pointer a: %v[%T]", &a, &a)
}
