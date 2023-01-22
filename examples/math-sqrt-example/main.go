package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	//Temp hold the real root of x
	var temp float64 = math.Sqrt(x)
	//z is the initial value of the root
	var z float64 = 1.0
	//The loop will continue when the root is equal to the real root
	for {
		z -= (z*z - x) / (2 * z)
		if temp == z {
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(6))
}
