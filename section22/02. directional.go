package main

import "fmt"

func main() {
	channel := make(chan int)
	sendChannel := make(<-chan int)
	reciveChannel := make(chan<- int)

	go func() {
		channel <- 44
		reciveChannel <- 22
	}()

	fmt.Printf("%T\n", channel)
	fmt.Printf("%T\n", sendChannel)
	fmt.Printf("%T\n", reciveChannel)
}
