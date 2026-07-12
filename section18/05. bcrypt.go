package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	bs, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)

	if err != nil {
		fmt.Println("Generate Password Error ", err)
	}

	fmt.Println("Generate Password:", bs)

	err = bcrypt.CompareHashAndPassword(bs, []byte("password1"))

	if err != nil {
		fmt.Println("Not correct Password")
		return
	}

	fmt.Println("Welcome ")
}
