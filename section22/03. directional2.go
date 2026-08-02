package main

import "fmt"

func main() {

	channel := make(chan int)

	go reciveValue(channel)

	sendValue(channel)

	fmt.Println("End Of main func")

}

func reciveValue(channel chan<- int) {
	channel <- 22
}

func sendValue(channel <-chan int) {
	fmt.Println(<-channel)
}
