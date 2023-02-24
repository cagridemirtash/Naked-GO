package pointers

import "fmt"

func PointerIntro(number *int) {
	*number += 1
	fmt.Println("Function number :", *number)
}
