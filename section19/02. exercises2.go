package main

import (
	"encoding/json"
	"fmt"
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
	s := `[{"First":"Lionel","Last":"Messi","Age":39},{"First":"Cristano","Last":"Ronaldo","Age":41}]`

	var people []person

	err := json.Unmarshal([]byte(s), &people)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	for _, person := range people {
		fmt.Println("------------------------")
		person.introduce()
	}
}
