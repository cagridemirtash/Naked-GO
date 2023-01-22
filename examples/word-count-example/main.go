package main

import (
	"strings"

	//This is the package that we are going to test
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mappingReturn := make(map[string]int)
	//Split the string into an array of strings
	tempArr := strings.Fields(s)
	//Iterate array and add to map
	for _, element := range tempArr {
		//If the key exists, increment the value
		if _, found := mappingReturn[element]; found {
			mappingReturn[element] += 1
			continue
		}
		//If the key does not exist, add it to the map with initial value 1
		mappingReturn[element] = 1
	}
	return mappingReturn
}

func main() {
	wc.Test(WordCount)
}
