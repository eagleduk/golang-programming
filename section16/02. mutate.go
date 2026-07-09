package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Printf("x before[main]: %v[%T]\n", x, x)
	foo(x)
	fmt.Printf("x after[main]: %v[%T]\n", x, x)

	fmt.Println("---------------")

	y := 99
	fmt.Printf("y before[main]: %v[%T]\n", y, y)
	bar(&y)
	fmt.Printf("y after[main]: %v[%T]\n", y, y)

}

func foo(x int) {
	fmt.Printf("x before[foo]: %v[%T]\n", x, x)
	x = 32
	fmt.Printf("x after[foo]: %v[%T]\n", x, x)
}

func bar(y *int) {
	fmt.Printf("y before[bar]: %v[%T]\n", y, y)
	*y = 32
	fmt.Printf("y after[bar]: %v[%T]\n", y, y)
}
