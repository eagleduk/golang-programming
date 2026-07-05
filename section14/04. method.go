package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func (a person) fullname() string {
	return a.firstname + " " + a.lastname
}

func main() {
	p1 := person{
		firstname: "James",
		lastname:  "Bond",
	}
	fullname := p1.fullname()

	fmt.Println(fullname)
}
