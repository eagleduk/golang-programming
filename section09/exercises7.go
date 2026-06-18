package main

import (
	"fmt"
)

func main() {

	x := 30

	if x < 30 {
		fmt.Println("'x' is less than 30.")
	} else if x == 30 {
		fmt.Println("'x' is 30.")
	} else {
		fmt.Println("'x' is greater than 30.")
	}
}
