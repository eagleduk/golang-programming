package main

import (
	"fmt"
)

func main() {

	// buffered channel
	c := make(chan int)
	c <- 44

	fmt.Println(<-c)
}
