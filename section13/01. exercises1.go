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

	fmt.Println("======= p1 ========")
	fmt.Printf("frist: %v, last: %v\n", p1.first, p1.last)
	for _, v := range p1.favorite {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()

	fmt.Println("======= p2 ========")
	fmt.Printf("frist: %v, last: %v\n", p2.first, p2.last)
	for _, v := range p2.favorite {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
}
