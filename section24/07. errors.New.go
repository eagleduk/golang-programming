package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var ErrorNew = errors.New("Error New:: ")

func main() {
	var a float64 = -10
	i, err := func_errorNew(a)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("a sqrt is", i)
	}

	fmt.Println("exit of main func")
}

func func_errorNew(i float64) (float64, error) {

	if i < 0 {
		return 0, ErrorNew
	}

	return math.Sqrt(i), nil
}
