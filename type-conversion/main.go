package main

import (
	"fmt"
)

func main() {
	var a int = 10
	// Classical way of basic type conversion -> Type(value)
	var b float64 = float64(a)
	fmt.Printf("a = %d, b = %f", a, b)
}
