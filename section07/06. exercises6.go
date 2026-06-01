package main

import (
	"fmt"
)

const year = 2026

const (
	a = year + iota + 1
	b = year + iota + 1
	c = year + iota + 1
	d = year + iota + 1
)

func main() {
	fmt.Println("This year: ", year)
	fmt.Printf("+1 %d, +2 %d, +3 %d, +4 %d", a, b, c, d)
}
