package main

import (
	"fmt"
)

func main() {
	c := exercises3_gen()
	exercises3_receive(c)

	fmt.Println("about to exit")
}

func exercises3_gen() <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c
}
