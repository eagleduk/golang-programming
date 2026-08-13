package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan string)

	go exercises7_put(c)

	for v := range c {
		fmt.Println("=======. ", v)
	}

	fmt.Println("exit of main")
}

func exercises7_put(c chan<- string) {

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for ii := 0; ii < 10; ii++ {
				c <- fmt.Sprintf("goroutine %v index %v", i, ii)
			}
			wg.Done()
			close(c)
		}()
	}
	wg.Wait()
}
