package main

import (
	"fmt"
	"log"
	"math"
)

type errorObject struct {
	text1 string
	text2 string
	err   error
}

func (t errorObject) Error() string {
	return fmt.Sprintf("Error::: %v %v ", t.text1, t.text2)
}

func main() {
	var a float64 = -10
	i, err := func_errorOf(a)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("a sqrt is", i)
	}

	fmt.Println("exit of main func")
}

func func_errorOf(i float64) (float64, error) {

	if i < 0 {
		n := fmt.Errorf("Error Of: ")
		return 0, errorObject{
			text1: "text1",
			text2: "text2",
			err:   n,
		}
	}

	return math.Sqrt(i), nil
}
