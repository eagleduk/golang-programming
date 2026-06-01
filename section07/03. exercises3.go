package main

import (
	"fmt"
)

const a = 34
const b = 3243.342
const c = "SDJIO"

var aa int = 39839
var bb float32 = 232.323
var cc string = "EFJIE"

func main() {
	fmt.Println("untyped ", a, b, c)

	fmt.Println("typed ", aa, bb, cc)
}
