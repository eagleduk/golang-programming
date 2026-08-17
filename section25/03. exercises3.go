package main

import (
	"fmt"
	"time"
)

type Exercises_ErrorObject struct {
	s   string
	err error
}

func (p Exercises_ErrorObject) Error() string {
	return fmt.Sprintf("[%v]: %v", time.Now, p.s)
}

func main() {
	v := Exercises_ErrorObject{
		s: "SSSSSSSS",
	}

	exercises_foo(v)
}

func exercises_foo(e error) {
	fmt.Println(e, " :::: ", e.(Exercises_ErrorObject).s)
}
