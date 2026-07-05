package main

import (
	"fmt"
)

func main() {
	foo()
	defer bar()
	foobar()

	fmt.Println("main")
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func foobar() {
	fmt.Println("foobar")
}
