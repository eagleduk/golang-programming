package main

import (
	"fmt"
)

func foo() {
	fmt.Println("This is foo method")
}

func main() {

	foo()

	func() {
		fmt.Println("This is anonymous method")
	}()

	func(x int) {
		fmt.Println("This is one parameter anonymous method:", x)
	}(44)
}
