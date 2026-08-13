package main

import (
	"fmt"
)

func main() {

	// buffered channel
	c := make(chan int, 1)
	c <- 44

	fmt.Println(<-c)
}
