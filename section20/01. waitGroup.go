package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("init func")
}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:\t\t\t\t", runtime.GOOS)
	fmt.Println("GO Arch: \t\t", runtime.GOARCH)
	fmt.Println("NumCPUs: \t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine: \t", runtime.NumGoroutine())

	wg.Add(3)

	go foo()
	go bar()

	fmt.Println("NumCPUs: \t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine: \t", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
	wg.Done()
}
