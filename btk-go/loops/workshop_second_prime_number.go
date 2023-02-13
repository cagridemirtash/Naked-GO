package loops

import "fmt"

func IsPrime() {
	var inputNumber int
	isPrime := true
	_, err := fmt.Scan(&inputNumber)
	if err != nil {
		fmt.Println("It is not a valid number")
		return
	}
	for i := 2; i < inputNumber; i++ {
		if inputNumber%i == 0 {
			isPrime = false
		}
	}
	if isPrime {
		fmt.Printf("%v number is prime.", inputNumber)
	} else {
		fmt.Printf("%v number is not prime.", inputNumber)
	}
}
