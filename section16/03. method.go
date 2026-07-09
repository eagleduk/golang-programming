package main

import (
	"fmt"
)

type text struct {
	text string
}

type P interface {
	print()
}

type PP interface {
	print2()
}

func (t text) print() {
	fmt.Println(t.text)
}

func (t *text) print2() {
	fmt.Println(t.text)
}

func pp(p P) {
	p.print()
}

func pp2(p PP) {
	p.print2()
}

func main() {
	a := text{
		text: "RES",
	}

	pp(a)
	pp(&a)

	pp2(&a)
	pp2(a) // cannot use a (variable of struct type text) as PP value in argument to pp2: text does not implement PP (method print2 has pointer receiver)

}
