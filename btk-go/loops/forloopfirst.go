package loops

import "fmt"

func ForLoopDemo() {
	var text string = "Hello, this statement execute with for loop"
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
}
