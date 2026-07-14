package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())

	counter := 0
	routine := 100

	var wg sync.WaitGroup

	wg.Add(routine)

	for i := 0; i < routine; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched() // go 루틴이 실행하면서 프로세서 양보(?)
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("GoRoutines", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("counter:", counter)
}
