package main

import (
	"fmt"
)

func main() {

	i := 33
	for i <= 122 {
		fmt.Printf("%v\t %x\t %#U\n", i, i, i)

		i++
	}
}
