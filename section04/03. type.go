package main

import "fmt"

type first_type int

var a first_type = 33
var b int = 44

func main() {

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	b = int(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
