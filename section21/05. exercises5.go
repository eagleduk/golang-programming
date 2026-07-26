package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64

	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {

			atomic.AddInt64(&counter, 1)

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("counter: ", counter)
}
