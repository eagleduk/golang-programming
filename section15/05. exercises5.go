package main

import (
	"fmt"
	"math"
)

type square struct {
	L float32
	W float32
}

type circle struct {
	r float32
}

func (s square) area() float32 {
	return s.L * s.W
}

func (c circle) area() float32 {
	return c.r * c.r * math.Pi
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println("shape: ", s.area())
}

func main() {
	s := square{
		L: 10.1,
		W: 12.2,
	}

	c := circle{
		r: 4.4,
	}

	info(s)
	info(c)
}
