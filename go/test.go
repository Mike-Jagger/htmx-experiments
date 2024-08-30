package main

import (
	"fmt"
	"math"
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

	fmt.Println()

	PrintTypeNValue(Bool)
	PrintTypeNValue(Int32)
	PrintTypeNValue(Uint)
	PrintTypeNValue(Float64)
	PrintTypeNValue(Rune)
	PrintTypeNValue(Byte)
	PrintTypeNValue(Complex64)

	fmt.Println()

	PrintTypeNValue(float64(Int))

	fmt.Println()

	integerValue := 24
	fmt.Println(math.Sqrt(2.22 + float64(integerValue)))

	fmt.Println()

	fmt.Printf("%T\n", 2)
	fmt.Printf("%T\n", 2.0)
	fmt.Printf("%T\n", 0 + 0i)

	fmt.Println()

	const cInt = 2
	const cString = "Constant string"
	const cBool = true
	const cChar = "世界"

	PrintTypeNValue(cInt)
	PrintTypeNValue(cChar)
	PrintTypeNValue(cString)
	PrintTypeNValue(cBool)

	fmt.Println()

	const (
		Big = 1 << 100
		Small = Big >> 99
	)

	fmt.Printf("This is going to return float even though constant \n%v \n", Big * 0.1)
	fmt.Printf("This is going to return int even though constant \n%v", Small * 10 + 1)




	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Printf("i is %v and sum is %v \n", i, sum)
		sum = i * 10
	} 


	var c bool = true
	
	for ; c; {
		if !c {
			fmt.Println("If it comes here then not while loop")
			break
		}

		fmt.Printf("Value of c is %v \n", c)

		c = false
	}
	fmt.Printf("Checking if condition is met and loop breaks, then while loop simulated \n")
	// var c, python = 1, true
	// java := "It is just Java bro, but with the dot:column declaration"

	// fmt.Println(add(1, 2))

	// a, b := swap("Hello", "World!")

	// fmt.Println(a, b)

	// fmt.Println(split(100))

	// fmt.Println(i, j, c, python, java)
}