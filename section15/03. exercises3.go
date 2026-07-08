package main

import (
	"fmt"
)

func main() {
	defer endup()

	startup()

	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
}

func startup() {
	fmt.Println("This func is start at main func start.")
}

func endup() {
	fmt.Println("This func is start at main func end.")
	defer func() {
		fmt.Println("This func is start at endup func end.")
	}()
}
