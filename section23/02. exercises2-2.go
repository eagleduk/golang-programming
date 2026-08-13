package main

import (
	"fmt"
)

func main() {

	cr := make(chan int)

	go func() {
		cr <- 42
	}()

	fmt.Println(<-cr)

	fmt.Println("--------")
	fmt.Printf("cr\t%T\n", cr)
}
