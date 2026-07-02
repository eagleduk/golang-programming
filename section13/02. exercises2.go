package main

import "fmt"

type person struct {
	first    string
	last     string
	favorite []string
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		favorite: []string{
			"choco", "banana",
		},
	}

	p2 := person{
		first: "A",
		last:  "BBB",
		favorite: []string{
			"QWER", "QWE",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for key, value := range m {

		fmt.Println("======= ", key, " ========")
		fmt.Printf("frist: %v, last: %v\n", value.first, value.last)
		for _, v := range value.favorite {
			fmt.Printf("%v\t", v)
		}
		fmt.Println()

	}

}
