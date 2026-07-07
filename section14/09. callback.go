package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := sum(ii...)
	fmt.Println("All number sum:", s)

	o := odd(sum, ii...)
	fmt.Println("odd number sum:", o)

	e := even(sum, ii...)
	fmt.Println("even number sum:", e)
}

func sum(i ...int) int {
	var total int
	for _, v := range i {
		total += v
	}
	return total
}

func even(f func(i ...int) int, e ...int) int {
	var ei []int

	for _, v := range e {
		if v%2 == 0 {
			ei = append(ei, v)
		}
	}
	return f(ei...)
}

func odd(f func(i ...int) int, e ...int) int {
	var oi []int

	for _, v := range e {
		if v%2 == 1 {
			oi = append(oi, v)
		}
	}
	return f(oi...)
}
