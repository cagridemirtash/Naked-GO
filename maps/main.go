package main

import "fmt"

type Name struct {
	FirstName string
	LastName  string
}

var nameMap map[string]Name

//The other version of using map
var nameMapOtherVersion = map[string]Name{
	"John":  Name{"John", "Doe"},
	"Çağrı": Name{"Çağrı", "Demirtaş"},
	"Samet": Name{"Samet", "Aydın"},
	"Galip": Name{"Galip", "Düğer"},
}

func main() {
	// Mapping to given a String Key and Name Value
	// When reach to a value, give to certain key of requested value
	fmt.Println("Maps")
	//Make function returns ready to use map for using with key-value pairs
	nameMap = make(map[string]Name)
	nameMap["John"] = Name{"John", "Doe"}
	nameMap["Çağrı"] = Name{"Çağrı", "Demirtaş"}
	nameMap["Samet"] = Name{"Samet", "Aydın"}
	fmt.Println(nameMap["Çağrı"].FirstName)
	fmt.Println(nameMapOtherVersion["Galip"].FirstName)
}
