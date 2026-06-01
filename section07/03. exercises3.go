package main

import (
	"fmt"
)

const (
	a = 34
	b = 3243.342
	c = "SDJIO"
)

var (
	aa int     = 39839
	bb float32 = 232.323
	cc string  = "EFJIE"
)

func main() {
	fmt.Println("untyped ", a, b, c)

	fmt.Println("typed ", aa, bb, cc)
}
