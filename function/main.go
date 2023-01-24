package main

import "fmt"

func main() {
	person := Cagri{"Çağrı", "Demirtaş"}
	firstNumber := Number{person.FirstName, 5}
	handlePersonWithNumber(person, firstNumber)
	// Call Receiver function
	person.handlePersonWithReceiver()
}

func handlePersonWithNumber(cagri Cagri, tempNumber Number) {
	if cagri.FirstName == tempNumber.key {
		fmt.Println("This is a relation between Number and Cagri struct.")
	}
}

// This function call Receiver Function
func (cagri Cagri) handlePersonWithReceiver() {
	fmt.Println("Welcome", cagri.FirstName)
}

type Cagri struct {
	FirstName  string
	SecondName string
}

type Number struct {
	key   string
	value int
}
