package main

import (
	"fmt"
	"golang-programming/section29/exercises2/quote"
	"golang-programming/section29/exercises2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
