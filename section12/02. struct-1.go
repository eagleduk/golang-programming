package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

type People struct {
	Person
	job string
	age int
}

func main() {

	people := People{
		Person: Person{
			first: "James",
			last:  "Bond",
			age:   11,
		},
		age: 44,
		job: "programer",
	}

	fmt.Println(people.first, people.last, people.age, people.Person.age)
}
