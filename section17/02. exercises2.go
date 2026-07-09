package main

import "fmt"

type person struct {
	first  string
	second string
}

func changeMe(person *person) {
	(*person).first = "CHANGE ME!!"
	person.second = "SSSS !!!"
}

func main() {
	p := person{
		first:  "First",
		second: "Second",
	}

	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)
}
