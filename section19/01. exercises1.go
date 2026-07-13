package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
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

	people := []person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(bs))
}
