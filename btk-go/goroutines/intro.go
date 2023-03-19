package goroutines

import (
	"fmt"
	"time"
)

func EvenNumbers() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("Even numbers :", i)
		time.Sleep(1 * time.Second)
	}
}

func OddNumbers() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("Odd numbers :", i)
		time.Sleep(1 * time.Second)
	}
}
