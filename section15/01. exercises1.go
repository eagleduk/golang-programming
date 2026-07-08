package main

import (
	"fmt"
)

func main() {

	a := foo()
	b, c := bar()

	fmt.Println("foo return:", a)
	fmt.Println("bar return:", b, c)
}

func foo() int {
	return 55
}

func bar() (int, string) {
	return 44, "string"
}
