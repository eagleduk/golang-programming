package main

import (
	"fmt"
)

func main() {

	a := 43
	b := "s"

	fmt.Printf("a - value[%T]: %v - address[%T]: %v\n", a, a, &a, &a)
	fmt.Printf("b - value[%T]: %v - address[%T]: %v\n", b, b, &b, &b)

	aa := &a
	bb := &b

	fmt.Printf("a - value[%T]: %v - address[%T]: %v\n", aa, aa, *aa, *aa)
	fmt.Printf("b - value[%T]: %v - address[%T]: %v\n", bb, bb, *bb, *bb)

	*aa = 11
	*bb = "c"

	fmt.Printf("a - value[%T]: %v - address[%T]: %v\n", a, a, &a, &a)
	fmt.Printf("b - value[%T]: %v - address[%T]: %v\n", b, b, &b, &b)
}
