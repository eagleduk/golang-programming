package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string `json:First`
	Last  string `json:Last`
	Age   int    `json:Age`
}

func (p person) introduce() {
	fmt.Println(p.First, p.Last, "[", p.Age, "]")
}

func main() {
	p1 := person{
		First: "Lionel",
		Last:  "Messi",
		Age:   39,
	}

	p2 := person{
		First: "Cristano",
		Last:  "Ronaldo",
		Age:   41,
	}

	p3 := person{
		First: "Wayne",
		Last:  "Rooney",
		Age:   41,
	}

	people := []person{p1, p2, p3}

	err := json.NewEncoder(os.Stdout).Encode(people)

	if err != nil {
		fmt.Println("Error", err)
	}
}
