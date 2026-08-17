package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := exercises_sqrt(-10.23)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("exit of main func")
}

func exercises_sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{
			lat:  "lat",
			long: "long",
			err:  errors.New(fmt.Sprintf("Less Than 0. %v", f)),
		}
	}
	return 42, nil
}
