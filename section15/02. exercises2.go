package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum1 := foo2(i...)

	sum2 := bar2(i)

	fmt.Println("sum1: ", sum1)
	fmt.Println("sum2: ", sum2)
}

func foo2(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func bar2(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
