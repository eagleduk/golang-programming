package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any Last wishes?", "Never say never"},
	}

	bs, err := exercises_toJSON(p1)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))

}

func exercises_toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return []byte{}, fmt.Errorf(fmt.Sprintf("Exercises 2: %v", err))
	}
	return bs, nil
}
