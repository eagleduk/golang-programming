package main

func section28_sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
