package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string `json:First`
	Last  string `json:Last`
	Age   int    `json:Age`
}

type sortAge []person

func (p sortAge) Len() int {
	return len(p)
}
func (p sortAge) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p sortAge) Less(i int, j int) bool {
	return p[i].Age > p[j].Age
}

type sortFirst []person

func (p sortFirst) Len() int {
	return len(p)
}
func (p sortFirst) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p sortFirst) Less(i int, j int) bool {
	return p[i].First > p[j].First
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

	fmt.Println("Before Sort", people)
	sort.Sort(sortAge(people))
	fmt.Println("Sort Age", people)

	sort.Sort(sortFirst(people))
	fmt.Println("Sort First", people)

}
