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

func main() {

	i := []int{1, 3, 4, 6, 8, 5, 3, 2, 5, 4}
	s := []string{"A", "R", "T", "D", "V", "B", "W"}

	fmt.Println("Before: ", i)
	sort.Ints(i)
	fmt.Println("After: ", i)

	fmt.Println("Before: ", s)
	sort.Strings(s)
	fmt.Println("After: ", s)

}
