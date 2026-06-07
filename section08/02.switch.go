package main

import (
	"fmt"
)

func main() {

	switch {
	case true:
		fmt.Println("true")
	case 2 == 4:
		fmt.Println("2==4")
	case 2 == 2:
		fmt.Println("2==2")
		fallthrough
	default:
		fmt.Println("default")
	}

	v := "B"

	switch v {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	case "C":
		fmt.Println("C")
		fallthrough
	case "D", "E":
		fmt.Println("D, E")
		fallthrough
	default:
		fmt.Println("default")
	}
}
