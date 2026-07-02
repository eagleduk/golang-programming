package main

import (
	"fmt"
)

func main() {

	a := struct {
		first  string
		second []string
		thrid  map[string]string
	}{
		first: "name",
		second: []string{
			"A", "B", "C",
		},
		thrid: map[string]string{
			"a": "AA",
			"b": "BB",
		},
	}

	fmt.Println(a)

}
