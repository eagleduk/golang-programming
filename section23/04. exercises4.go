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

func exercises4_gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 0
	}()

	return c
}

func exercises4_receive(c, q <-chan int) {

	for {
		select {
		case v := <-c:
			fmt.Println("Channel c", v)
		case v := <-q:
			fmt.Println("Channel q", v)
			return
		}
	}
}
