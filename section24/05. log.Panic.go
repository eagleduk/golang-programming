package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.ReadFile("OS.txt")

	if err != nil {
		log.Panic("log.Panic ", err)
	}
}
