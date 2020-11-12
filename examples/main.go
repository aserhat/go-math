package main

import (
	"fmt"

	gomath "github.com/aserhat/go-math"
)

func main() {
	i := &gomath.Inputs{
		InputA: 10,
		InputB: 5,
	}

	fmt.Println("Lets go do some math!")
	fmt.Println(i.Add())
	fmt.Println(i.Subtract())
	fmt.Println(i.Multiply())
	fmt.Println(i.Divide())
}
