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
}
