package main

import "fmt"

func main() {
	//Declaration with var keyword and limited size
	var myArr [2]int
	myArr[0] = 1
	myArr[1] = 2
	//Other Declaration of array
	stringArr := []string{
		"Çağrı",
		"Çiko",
		"Çağan",
	}
	//If we need to index of Array -> This for statement is fitting our situation
	for index, value := range myArr {
		fmt.Println(index, ":", value)
	}
	//If we need just value an array -> These for is fitting
	for _, value := range myArr {
		fmt.Println(value)
	}
	for _, value := range stringArr {
		fmt.Println(value)
	}
}
