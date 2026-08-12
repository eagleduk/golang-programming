package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	c1 := make(chan int)
	r := make(chan string)

	go func() {
		for i := 0; i < 100; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go fanOut(c1, r)

	for s := range r {
		fmt.Println(s)
	}

	fmt.Println("End Of Func.")
}

func fanOut(c1 <-chan int, r chan<- string) {

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					r <- fmt.Sprintf("index[%v]: %v", v2, rand.Intn(10))
				}(v)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(r)
}
