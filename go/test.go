package main

import (
	"fmt"
	"math/cmplx"
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

// var i, j int = 1, 2

var (
	Bool bool = true

	String string = "string variable"
	
	Int int = 69
	Int8 int8
	Int16 int16
	Int32 int32
	Int64 int64

	Uint uint
	Uint8 uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64 = 1<<64 - 1
	Uintptr uintptr

	Byte byte

	Rune rune

	FLoat32 float32
	Float64 float64

	Complex64 complex64
	Complex128 complex128 = cmplx.Sqrt(-5 + 12i)
)

func PrintTypeNValue(_var interface{}) {
	fmt.Printf("Type: %T Value: %v \n", _var, _var)
}

func main() {
	PrintTypeNValue(Bool)
	PrintTypeNValue(String)
	PrintTypeNValue(Int)
	PrintTypeNValue(Uint64)
	PrintTypeNValue(Complex128)


	// var c, python = 1, true
	// java := "It is just Java bro, but with the dot:column declaration"

	// fmt.Println(add(1, 2))

	// a, b := swap("Hello", "World!")

	// fmt.Println(a, b)

	// fmt.Println(split(100))

	// fmt.Println(i, j, c, python, java)
}