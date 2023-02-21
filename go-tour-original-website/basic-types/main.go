package main

import (
	"fmt"
)

/*
!Data Types
* string
* int -> int8, int16, int32, int64
* uint -> uint8, uint16, uint32, uint64, uintptr
* byte -> alias for uint8
* rune -> alias for int32
* float32, float64
* complex64, complex128
*/

/*
* Variables when the initial value is not assigned, the default value is assigned.
* For Example: int -> 0, float -> 0.0, string -> "", bool -> false
 */

var name string = "Çağrı Demirtaş"
var age, height int64 = 23, 189
var isBadSide bool = true
var cagri string
var first, second, third int

func main() {
	fmt.Println("Your Name :", name)
	fmt.Println("Your Age :", age, "Your Height :", height)
	// To initialize values which are cagri => "", first, second, third are 0
	// Assign to default value from GO.
	fmt.Println(cagri, first, second, third)
	// !If Condition -> We can use with assign a variable and use it
	if condition := isBadSide; condition {
		fmt.Println("You are a bad side")
		return
	}
	fmt.Println("You are a good side")
}
