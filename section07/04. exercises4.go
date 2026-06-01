package main

import (
	"fmt"
)

const a = 44

func main() {
	fmt.Printf("%d %b %x\n", a, a, a)

	var b = a << 1

	fmt.Printf("%d %b %x", b, b, b)
}
