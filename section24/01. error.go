package main

import (
	"fmt"
)

func main() {
	_, err := fmt.Println("Hello")

	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println("exit of main func")
}
