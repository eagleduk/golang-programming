package main

import (
	"encoding/json"
	"fmt"
)

type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
}

func main() {

	json_string := `[{"First":"James","Last":"Bond"},{"First":"Lionel","Last":"Messi"},{"First":"Cristiano","Last":"Ronaldo"}]`
	json_byte := []byte(json_string)

	var people []person2

	err := json.Unmarshal(json_byte, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, v := range people {
		fmt.Println(i, v)
	}

}
