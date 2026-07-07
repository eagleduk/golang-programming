package main

import (
	"fmt"
)

type cat struct {
	name string
	jump int
}

type dog struct {
	name string
	run  int
}

func (c cat) say() {
	fmt.Println("meow")
}

func (d dog) say() {
	fmt.Println("brak")
}

type saying interface {
	say()
}

func animalSay(s saying) {
	switch s.(type) {
	case cat:
		fmt.Printf("My name is %v, i can jump %v miters.\n", s.(cat).name, s.(cat).jump)
	case dog:
		fmt.Printf("My name is %v, i can run %v m/s.\n", s.(dog).name, s.(dog).run)
	}
}

func main() {

	navi := cat{
		name: "navi",
		jump: 15,
	}

	tomas := dog{
		name: "tomas",
		run:  20,
	}

	navi.say()
	tomas.say()

	animalSay(navi)
	animalSay(tomas)

}
