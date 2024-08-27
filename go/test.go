package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

func main() {
	var c, python = 1, true
	java := "It is just Java bro, but with the dot:column declaration"

	fmt.Println(add(1, 2))
	
	a, b := swap("Hello", "World!")

	fmt.Println(a, b)

	fmt.Println(split(100))

	fmt.Println(i, j, c, python, java)
}