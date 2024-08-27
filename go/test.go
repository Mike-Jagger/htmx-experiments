package main

import (
	"fmt"
	"math"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("My favorite number is: %g", math.Sqrt(8))
	fmt.Println(quote.Go())
}