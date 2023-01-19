package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	a, b:= show();
	fmt.Println(a + b)
}

func show() (x,y int) {
	x = 42;
	y = 43;
	return 
}
