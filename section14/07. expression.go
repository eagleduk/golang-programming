package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("This is expression function")
	}

	f()

	g := func(x int) {
		fmt.Printf("This is %v years.\n", x)
	}

	g(2026)
}
