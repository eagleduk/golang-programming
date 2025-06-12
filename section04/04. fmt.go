package main

import "fmt"

func main() {
	n, e := fmt.Println("Hello World")
	fmt.Println("i: ", n, ", e: ", e)

	n, _ = fmt.Println("My First Go Programing")
	fmt.Println("i: ", n)

	v := "H"
	a := fmt.Sprintf("variable 'v' type is %T", v)

	fmt.Println(a)

}
