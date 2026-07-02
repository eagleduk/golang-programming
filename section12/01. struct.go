package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := Person{
		first: "James",
		last:  "Bond",
		age:   11,
	}

	fmt.Println(p1.first, p1.last, p1.age)
}
