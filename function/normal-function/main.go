package main

import "fmt"

func main() {
	person := Person{"Çağrı", "Demirtaş"}
	firstNumber := Number{person.FirstName, 5}
	handlePersonWithNumber(person, firstNumber)
	// Call Receiver function
	person.handlePersonWithReceiver()
}

func handlePersonWithNumber(person Person, tempNumber Number) {
	if person.FirstName == tempNumber.key {
		fmt.Println("This is a relation between Number and Cagri struct.")
	}
}

// This function call Receiver Function
func (person Person) handlePersonWithReceiver() {
	fmt.Println("Welcome", person.FirstName)
}

type Person struct {
	FirstName  string
	SecondName string
}

type Number struct {
	key   string
	value int
}
