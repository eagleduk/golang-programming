package main

import (
	"fmt"
	"sync"
)

func main() {

	counter := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()

			v := counter
			v++
			counter = v
			fmt.Println(" ", counter)

			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("counter: ", counter)
}
