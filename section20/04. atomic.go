package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())

	var counter int64
	routine := 100

	var wg sync.WaitGroup

	wg.Add(routine)

	for i := 0; i < routine; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched() // go 루틴이 실행하면서 프로세서 양보(?)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("GoRoutines", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("counter:", counter)
}
