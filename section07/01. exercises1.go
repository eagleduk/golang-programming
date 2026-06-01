package main

import (
	"fmt"
)

var a = 33

func main() {
	fmt.Println("decimal ", a)
	fmt.Printf("binary %b\n", a)
	fmt.Printf("binary %#b\n", a)
	fmt.Printf("hex %x\n", a)
	fmt.Printf("hex %#x\n", a)
}
