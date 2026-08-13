package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := exercises4_gen(q)

	exercises4_receive(c, q)

	fmt.Println("about to exit")
}

func exercises4_gen(q <-chan int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}
	return c
}
