package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	if v, ok := m["James"]; ok {
		fmt.Println("is in key? ", v)
	}

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("is in key? ", v)
	}

	m["todd"] = 33
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, " - ", v)
	}
}
