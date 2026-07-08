package main

import (
	"fmt"
)

func main() {

	c1 := closure(10)
	c2 := closure(0)
	fmt.Println("run[c1] 1:", c1())
	fmt.Println("run[c1] 2:", c1())
	fmt.Println("run[c2] 1:", c2())
	fmt.Println("run[c1] 3:", c1())
	fmt.Println("run[c1] 4:", c1())
	fmt.Println("run[c2] 2:", c2())
	fmt.Println("run[c2] 3:", c2())
	fmt.Println("run[c2] 4:", c2())

}

func closure(i int) func() int {
	return func() int {
		i += 1
		return i
	}
}
