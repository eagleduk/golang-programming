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

	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T", y)
}
