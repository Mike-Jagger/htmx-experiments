package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"time"
)

// func add(x, y int) int {
// 	return x + y
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

// var i, j int = 1, 2


func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

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

func Sqrt(x float64) float64 {
	var z float64 = float64(1)
	
	var prev float64
	
	for ; z > math.Sqrt(x) || prev != z; {
		fmt.Printf("z:%v x:%v \n", z, math.Sqrt(x))
		z -= (z*z - x) / (2 * z)
		
		prev = z
	}
	
	return z
}

func main() {
	// fmt.Println(Sqrt(32))

	fmt.Print("Go runs on: ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	saturday := time.Saturday

	switch today := time.Now().Weekday(); saturday { 
	case today:
		fmt.Println("It's Saturday!")
	case today + 1:
		fmt.Println("It's tomorrow")
	case today + 2:
		fmt.Println("It's in 2 days")
	default:
		fmt.Println("It is too far away")
	}

	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Print("Good evening")
	}
	// fmt.Println(
	// 	pow(3, 2, 9),
	// 	pow(2, 2, 10),
	// )


	// PrintTypeNValue(Bool)
	// PrintTypeNValue(String)
	// PrintTypeNValue(Int)
	// PrintTypeNValue(Uint64)
	// PrintTypeNValue(Complex128)

	// fmt.Println()

	// PrintTypeNValue(Bool)
	// PrintTypeNValue(Int32)
	// PrintTypeNValue(Uint)
	// PrintTypeNValue(Float64)
	// PrintTypeNValue(Rune)
	// PrintTypeNValue(Byte)
	// PrintTypeNValue(Complex64)

	// fmt.Println()

	// PrintTypeNValue(float64(Int))

	// fmt.Println()

	// integerValue := 24
	// fmt.Println(math.Sqrt(2.22 + float64(integerValue)))

	// fmt.Println()

	// fmt.Printf("%T\n", 2)
	// fmt.Printf("%T\n", 2.0)
	// fmt.Printf("%T\n", 0 + 0i)

	// fmt.Println()

	// const cInt = 2
	// const cString = "Constant string"
	// const cBool = true
	// const cChar = "世界"

	// PrintTypeNValue(cInt)
	// PrintTypeNValue(cChar)
	// PrintTypeNValue(cString)
	// PrintTypeNValue(cBool)

	// fmt.Println()

	// const (
	// 	Big = 1 << 100
	// 	Small = Big >> 99
	// )

	// fmt.Printf("This is going to return float even though constant \n%v \n", Big * 0.1)
	// fmt.Printf("This is going to return int even though constant \n%v", Small * 10 + 1)




	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("i is %v and sum is %v \n", i, sum)
	// 	sum = i * 10
	// } 


	// var c bool = true
	
	// for c {
	// 	if !c {
	// 		fmt.Println("If it comes here then not while loop")
	// 		break
	// 	}

	// 	fmt.Printf("Value of c is %v \n", c)

	// 	c = false
	// }
	
	// fmt.Printf("Checking if condition is met and loop breaks, then while loop simulated \n")
	

	// fmt.Println(sqrt(2), sqrt(-4))
	// var c, python = 1, true
	// java := "It is just Java bro, but with the dot:column declaration"

	// fmt.Println(add(1, 2))

	// a, b := swap("Hello", "World!")

	// fmt.Println(a, b)

	// fmt.Println(split(100))

	// fmt.Println(i, j, c, python, java)
}