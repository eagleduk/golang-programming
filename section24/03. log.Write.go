package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	lf, err := os.Create("err.log")
	if err != nil {
		fmt.Println(err)
	}
	defer lf.Close()
	log.SetOutput(lf)

	lf2, err := os.Open("text.txt")
	if err != nil {
		log.Println("Not Found File")
	}

	lf2.Close()

}
