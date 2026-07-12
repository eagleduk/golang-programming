package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
	}

	p2 := person{
		First: "Lionel",
		Last:  "Messi",
	}

	p3 := person{
		First: "Cristiano",
		Last:  "Ronaldo",
	}

	people := []person{p1, p2, p3}

	json, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json))
}
