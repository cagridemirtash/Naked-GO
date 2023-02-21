package maps

import "fmt"

func Intro() {
	//Make a Dictionary which uses for translate word from English to German
	translator := make(map[string]string)
	translator["table"] = "tisch"
	translator["kid"] = "kind"
	translator["telephone"] = "telefon"
	translator["young"] = "jugendlich"

	PrintMap(translator)
	fmt.Println()
	//Delete Operation on maps -> delete(mapName, key without map use case)
	delete(translator, "table")
	PrintMap(translator)

	value, isThere := translator["kid"]
	if isThere {
		fmt.Println("There is in dictionary. Value is :", value)
	} else {
		fmt.Println("There is not in dictionary")
	}
}

func PrintMap(translator map[string]string) {
	for i, v := range translator {
		fmt.Printf("English : %s -> German: %s \n", i, v)
	}
}
