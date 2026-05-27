package main

import (
	"fmt"
)

type abc int

var x abc

func main() {
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
