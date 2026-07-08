package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Println("This is anonymous func1.")
	}()

	a := func() {
		fmt.Println("This is an anonymous func2, and variables.")
	}

	a()
}
