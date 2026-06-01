package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	aa = iota
	bb = iota
	cc = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
}
