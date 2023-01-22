package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	var person Person
	person.Name = "John"
	person.Surname = "Doe"
	person.Age = 30

	var person1 Person
	person1.Name = "Çağrı"
	person1.Surname = "Demirtaş"
	person1.Age = 23

	var person2 Person
	person2.Name = "Samet"
	person2.Surname = "Aydın"
	person2.Age = 25

	var persons [3]Person
	persons[0] = person
	persons[1] = person1
	persons[2] = person2

	for _, element := range persons {
		fmt.Println(element.Name)
	}
}
