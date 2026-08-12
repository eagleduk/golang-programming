package main

import (
	"fmt"
)

func main() {
	v := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			v <- i
		}
		close(v)
	}()

	for i := range v {
		fmt.Println(i)
	}

	fmt.Println("End of Main Func")

}
