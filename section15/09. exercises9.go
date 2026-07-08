package main

import (
	"fmt"
)

func main() {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	r := set1(sum, v...)

	fmt.Println("value[v] sum: ", r)

}

func sum(i []int) int {
	sum := 0

	for _, v := range i {
		sum += v
	}
	return sum
}

func set1(callback func(ii []int) int, i ...int) int {
	return callback(i)
}
