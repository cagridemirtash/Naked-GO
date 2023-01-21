package main

import (
	"fmt"
)

func main() {
	var isPrime bool = true
	var primeNumber int = 21
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
}
