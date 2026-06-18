package main

import (
	"fmt"
)

func main() {

	x := 30

	switch {
	case x < 30:
		fmt.Println("'x' is less than 30.")
	case x == 30:
		fmt.Println("'x' is 30.")
	case 30 < x:
		fmt.Println("'x' is greater than 30.")
	}
}
