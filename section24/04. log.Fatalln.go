package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	_, err := os.ReadFile("OS.txt")
	defer fatallnfunc()

	if err != nil {
		fatalln(err)
	}

	fmt.Println("exit of main func")
}

func fatalln(err error) {
	log.Fatalln("fatalln error:", err)
}

func fatallnfunc() {
	fmt.Println("defer fatalln function")
}
