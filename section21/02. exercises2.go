package main

import "fmt"

type Person struct {
	first string
	last  string
}

type hello interface {
	helloPrint()
}

func (p Person) helloPrint() {
	fullname := p.first + " " + p.last
	fmt.Println("Hello, ", fullname)
}

type helloo interface {
	hellooPrint()
}

func (p *Person) hellooPrint() {
	fullname := p.first + " " + p.last
	fmt.Println("Helloo, ", fullname)
}

func callHello(h hello) {
	h.helloPrint()
}

func callHelloPointer(h helloo) {
	(h).hellooPrint()
}

func main() {

	p1 := Person{
		first: "Lionel",
		last:  "Messi",
	}
	callHello(p1)

	p2 := Person{
		first: "Cristiano",
		last:  "Ronaldo",
	}
	callHello(&p2)

	p3 := Person{
		first: "Wayne",
		last:  "Rooney",
	}
	callHelloPointer(&p3)

	p4 := Person{
		first: "Park",
		last:  "JiSung",
	}
	p4.helloPrint()
	p4.hellooPrint()
	// callHelloPointer(p4)
}
