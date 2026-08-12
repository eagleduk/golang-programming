package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	r := make(chan string)

	go putin(c1)

	go putin(c2)

	go fanin(c1, c2, r)

	for v := range r {
		fmt.Println(v)
	}

	fmt.Println("End Of FanIn")

}

func fanin(c1, c2 <-chan int, r chan<- string) {
	for i := 0; i < 20; i++ {
		select {
		case v := <-c1:
			r <- fmt.Sprintf("Value Of C1: %d", v)
		case v := <-c2:
			r <- fmt.Sprintf("Value Of C2: %d", v)
		}
	}
	close(r)
}

func putin(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
