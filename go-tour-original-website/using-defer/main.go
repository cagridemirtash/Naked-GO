package main

import "fmt"

func main() {
	//Defer keyword using for wait until function execution is finished
	//and then execute the defer function
	//If the process declare with defer, send to stack and wait until the function is finished
	defer fmt.Println("world")
	fmt.Println("hello")
	//For example, in this section, counting and done see on console, but the numbers are reversed end of the console. Also, "world" is done after printing numbers. So this principle refer to last in first out -> LIFO
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	// When Change the value of variable, it cannot change because defer process push to memory
	var a int = 5
	fmt.Println("The first value of a", a)
	defer fmt.Println("The last value of a", a)
	a = 7

	//Same situation in the pointer, cannot change.
	var b int = 5
	d := &b
	//If d pointer change this line, b had a value like 7. The structure is very important.
	fmt.Println("The first value of b", b)
	defer fmt.Println("The last value of b", b)
	*d = 7
	fmt.Println("The first value of b", b)
}
