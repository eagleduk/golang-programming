package main

import (
	"fmt"
)

type abc int

var x abc
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)

	x := "James Bond"

	y = x

	fmt.Println(y)
	fmt.Printf("%T", y)
}
