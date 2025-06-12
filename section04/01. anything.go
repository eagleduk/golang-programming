package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")

	// loop; interative
	for i := 0; i < 100; i++ {
		// conditional
		if i%2 == 0 {
			fmt.Println("i: ", i)
		}
	}
}
