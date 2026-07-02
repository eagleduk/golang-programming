package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "grey",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Printf("doors: %v \n color: %v \n fourWheel: %v \n", t.doors, t.color, t.fourWheel)

	fmt.Println(s)
	fmt.Printf("doors: %v \n color: %v \n luxury: %v \n", s.doors, s.color, s.luxury)
}
