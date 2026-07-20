package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:\t\t\t\t", runtime.GOOS)
	fmt.Println("GO Arch: \t\t", runtime.GOARCH)
	fmt.Println("NumCPUs: \t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine: \t", runtime.NumGoroutine())

	wg.Add(2)

	go foo()
	go bar()

	fmt.Println("NumCPUs: \t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine: \t", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 14; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 11; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}
