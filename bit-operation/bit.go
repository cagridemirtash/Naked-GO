package main

import (
	"fmt"
)

func main() {
	//Shift left 10 bit -> 1^10 = 1024
	b := 1 << 10
	//Shift rigth 1 bit -> 6144 -> 0100 0000 0000 -> 1024 * 6 -> 6144
	//So shift 1 bit -> 0010 0000 0000 -> 3072 -> 6144 / 2
	a := 6144 >> 1
	fmt.Println("a = ", a)
	fmt.Printf("a is type of '%T' \n", a)
	fmt.Println("b = ", b)
}
