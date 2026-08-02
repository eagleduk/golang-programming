package main

import "fmt"

func ErrorMakeChannel() {
	channel := make(chan int)

	channel <- 44

	fmt.Println(<-channel)
}

func PutValueInGoroutine1() {
	channel := make(chan int)

	go func() {
		channel <- 44
	}()

	fmt.Println(<-channel)
}

func PutValueInGoroutine2() {
	channel := make(chan int, 1)

	channel <- 44

	fmt.Println(<-channel)
}

func PutMultiValueInGoroutine() {
	channel := make(chan int, 2)

	channel <- 44
	channel <- 45

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
func main() {
	PutValueInGoroutine1()

	PutValueInGoroutine2()

	PutMultiValueInGoroutine()
}
