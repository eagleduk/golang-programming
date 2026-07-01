package main

import "fmt"

func main() {
	a := [][]string{}

	a = append(a, []string{"James", "Bond", "Shaken, not stirred"}, []string{"Miss", "Moneypenny", "Helloooooo, James"})

	fmt.Println(a)

	for i, v := range a {
		for ii, vv := range v {
			fmt.Println("[", i, "][", ii, "] == ", vv)
		}
	}
}
