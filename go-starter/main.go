package main

import "fmt"

var c, python, java bool
var l,j,k int = 1,2,3
func main() {
	fmt.Println("Hello, world!")
	//The implicit type of variable (:=) -> var [variable name] type 
	a, b:= show();
	fmt.Println(a + b)

	//The Using Bool Type
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println("l j k :",l, j, k)
}

func show() (x,y int) {
	x = 42;
	y = 43;
	return 
}
