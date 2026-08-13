package main

import (
	"fmt"
)

func main() {

	// anonymous func
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
