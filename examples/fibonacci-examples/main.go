package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// Return as like "0, 1, 1, 2, 3, 5, 8, 13, 21, ..."
func fibonacci() func() int {
	//Initialize the value first and second element of fibonacci
	first := 0
	second := 1
	return func() int {
		//Return value is the first value which is refers to current value of fibonacci
		last := first
		//Assign second value to first value for using other iteration
		first = second
		//The same process but second element is that sum of first and second element
		//So that, last variable hold the value of first element.
		//Add last and first
		second = last + first
		return last
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
