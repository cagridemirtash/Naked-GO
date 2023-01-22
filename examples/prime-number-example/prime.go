package main

import (
	"fmt"
)

func main() {
	var isPrime bool = true
	var primeNumber int = 21

	/* Other Form of using for
		for ; sum < 1000; {
		sum += sum
		fmt.Println(sum)
	}
	*/

	for i := 2; i < primeNumber; i++ {
		if primeNumber%i == 0 {
			isPrime = false
		}
	}
	if isPrime {
		fmt.Println("Prime number")
	} else {
		fmt.Println("Not prime number")
	}

	// So how to use while loop in go

	/* Drop the semi colons and write for instead of while
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	*/
}
