package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("fullname: ", p.first, p.last, " age: ", p.age)
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   44,
	}

	p.speak()
}
