package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.ReadFile("OS.txt")

	if err != nil {
		log.Println("log.Println ", err)
	}
}
