package main

import (
	"fmt"
	"sort"
)

func main() {

	ii := []int{1, 9, 8, 7, 5, 34, 7, 4, 6, 8, 3}
	ss := []string{"a", "f", "t", "k", "v", "m"}

	fmt.Println(ii)
	sort.Ints(ii)
	fmt.Println(ii)

	fmt.Println(" ----------- ")

	fmt.Println(ss)
	sort.Strings(ss)
	fmt.Println(ss)
}
