package main

import (
	"fmt"
)

func main() {
	//Scanln which is hold the value of strings
	fmt.Println("Please enter your First Name")
	var firstName string
	_, err := fmt.Scanln(&firstName)
	if err != nil {
		return
	}
	fmt.Println("Please enter your Second Name")
	var secondName string
	_, err2 := fmt.Scanln(&secondName)
	if err2 != nil {
		return
	}
	fmt.Println("Your Full Name is :", firstName, secondName)
}
