package main

import (
	"fmt"
)

func main() {

	sum_f := set(10, 20)

	sum := sum_f()

	fmt.Println("sum:", sum)

}

func set(a int, b int) func() int {
	return func() int {
		return a + b
	}
}
