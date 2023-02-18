package functions

import "fmt"

func Add(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}

func SayHi(userName string) {
	fmt.Println("Hello", userName)
}
