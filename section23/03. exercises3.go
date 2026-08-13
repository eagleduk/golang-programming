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

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func exercises3_receive(c <-chan int) {

	for v := range c {
		fmt.Println(v)
	}

}
