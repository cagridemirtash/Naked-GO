package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var c, python, java bool
var l, j, k int = 1, 2, 3

func main() {
	fmt.Println("Hello, world!")
	//The implicit type of variable (:=) -> var [variable name] type
	a, b := show()
	fmt.Println(a + b)

	//The Using Bool Type
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println("l j k :", l, j, k)

	//Mathematical Operations
	fmt.Println("Mathematical Operations")
	fmt.Println("Sum :", math.Abs(42+13))
	fmt.Println("Subtraction :", math.Abs(5-13))

	//The Using Complex Type -> example of Java printf function.
	//%T -> Type of variable
	//%v -> Value of variable
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Naked Return
func show() (x, y int) {
	x = 42
	y = 43
	return
}
