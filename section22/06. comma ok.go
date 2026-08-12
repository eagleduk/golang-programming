package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		c1 <- 11
		c2 <- 33

		close(c2)
	}()

	cc1, ok := <-c1
	fmt.Println("Not Close Channel: ", cc1, ok)

	cc2, ok := <-c2
	fmt.Println("Close Channel: ", cc2, ok)
	cc2, ok = <-c2
	fmt.Println("Close Channel: ", cc2, ok)
}
